package serve

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/internal/attendances"
	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/SocBongDev/soc-bong/internal/registrations"
	"github.com/gofiber/fiber/v2"
	"github.com/pocketbase/dbx"

	_ "github.com/SocBongDev/soc-bong/docs"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type App struct {
	config *config.Config
}

func healthz(c *fiber.Ctx) error {
	return c.JSON(getHealthCheckResonse())
}

func index(c *fiber.Ctx) error {
	return c.Redirect("/docs")
}

func (a *App) RegisterAPIHandlers(router fiber.Router, handlers []common.APIHandler) {
	for _, handler := range handlers {
		handler.RegisterRoute(router)
	}
}

func (a *App) ApiV1(api fiber.Router, db *dbx.DB) {
	v1 := api.Group("/v1")
	publicHandlers := []common.APIHandler{}
	a.RegisterAPIHandlers(v1, publicHandlers)

	v1.Use(middlewares.ValidateJWT(a.config.Audience, a.config.Domain))

	privateHandlers := []common.APIHandler{
		attendances.New(attendances.NewRepo(db)),
		classes.New(classes.NewRepo(db)),
		registrations.New(registrations.NewRepo(db)),
	}
	a.RegisterAPIHandlers(v1, privateHandlers)
}

func (a *App) RunHttpServer() {
	db, err := database.New(&a.config.DatabaseSecret)
	if err != nil {
		log.Panic("Error create dbx: ", err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/healthz", healthz)
	app.Get("/", index)

	api := app.Group("/api")
	a.ApiV1(api, db)

	app.Get(
		"/api/messages/admin",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:admin-messages"}),
		func(c *fiber.Ctx) error {
			return c.JSON(map[string]string{"message": "Lmao"})
		},
	)

	go func() {
		if err := app.Listen(":5000"); err != nil {
			log.Panicln("App.Listen err: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGTERM,
	)

	_ = <-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg}
}
