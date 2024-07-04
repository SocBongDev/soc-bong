package serve

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/internal/agencies"
	"github.com/SocBongDev/soc-bong/internal/attendances"
	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/SocBongDev/soc-bong/internal/registrations"
	"github.com/SocBongDev/soc-bong/internal/spreadsheet"
	"github.com/SocBongDev/soc-bong/internal/students"
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
	agencyRepo, attendanceRepo, classRepo, registrationRepo, studentRepo := agencies.NewRepo(
		db,
	), attendances.NewRepo(
		db,
	), classes.NewRepo(
		db,
	), registrations.NewRepo(
		db,
	), students.NewRepo(
		db,
	)
	spreadsheet := spreadsheet.New()

	publicHandlers := []common.APIHandler{}
	a.RegisterAPIHandlers(v1, publicHandlers)

	v1.Use(middlewares.ValidateJWT(a.config.Audience, a.config.Domain))

	privateHandlers := []common.APIHandler{
		agencies.New(agencyRepo),
		attendances.New(attendanceRepo, classRepo, spreadsheet, studentRepo),
		classes.New(classRepo),
		registrations.New(registrationRepo),
		students.New(studentRepo),
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
	//registrations
	app.Get(
		"/api/v1/registrations",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read registrations"})
		},
	)

	app.Get(
		"/api/v1/registrations/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read one registration"})
		},
	)

	app.Post(
		"/api/v1/registrations",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"create:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "create a registration"})
		},
	)

	app.Patch(
		"/api/v1/registrations/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update one registration"})
		},
	)

	app.Put(
		"/api/v1/registrations/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update one registration"})
		},
	)

	app.Delete(
		"/api/v1/registrations/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"delete:registrations"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "delete one registration"})
		},
	)
	// agencies
	app.Get(
		"/api/v1/agencies",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read agencies"})
		},
	)

	app.Get(
		"/api/v1/agencies/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read one agency"})
		},
	)

	app.Post(
		"/api/v1/agencies",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"create:agencies"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "create an agency"})
		},
	)

	app.Put(
		"/api/v1/agencies/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:agencies"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update an agency"})
		},
	)
	//attendances
	app.Get(
		"/api/v1/attendances",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:attendances"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read attendances"})
		},
	)

	app.Post(
		"/api/v1/attendances",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"create:attendances"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "create attendances"})
		},
	)

	app.Patch(
		"/api/v1/attendances",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:attendances"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update attendances"})
		},
	)

	app.Get(
		"/api/v1/attendances/:classId<int,min(1)>/export-excel",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:attendances"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read attendances excel"})
		},
	)

	//classes
	app.Get(
		"/api/v1/classes",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:classes"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read classes"})
		},
	)

	app.Get(
		"/api/v1/classes/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:classes"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read one class"})
		},
	)

	app.Post(
		"/api/v1/classes",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"create:classes"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "create a class"})
		},
	)

	app.Put(
		"/api/v1/classes/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:classes"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update one class"})
		},
	)

	//students
	app.Get(
		"/api/v1/students",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:students"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read students"})
		},
	)

	app.Get(
		"/api/v1/students/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"read:students"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "read one student"})
		},
	)

	app.Post(
		"/api/v1/students",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"create:students"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "create new students"})
		},
	)

	app.Put(
		"/api/v1/students/:id<int,min(1)>",
		middlewares.ValidateJWT(a.config.Audience, a.config.Domain),
		middlewares.ValidatePermissions([]string{"update:students"}),
		func(c *fiber.Ctx) error {
			log.Print("err", err)
			return c.JSON(map[string]string{"message": "update one student"})
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
