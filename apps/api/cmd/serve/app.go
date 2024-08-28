package serve

import (
	_ "github.com/SocBongDev/soc-bong/docs"
	"github.com/SocBongDev/soc-bong/internal/agencies"
	"github.com/SocBongDev/soc-bong/internal/attendances"
	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/SocBongDev/soc-bong/internal/registrations"
	"github.com/SocBongDev/soc-bong/internal/roles"
	"github.com/SocBongDev/soc-bong/internal/spreadsheet"
	"github.com/SocBongDev/soc-bong/internal/students"
	"github.com/SocBongDev/soc-bong/internal/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	mdwlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/pocketbase/dbx"
)

type App struct {
	app    *fiber.App
	config *config.Config
	db     *dbx.DB
}

func NewApp(cfg *config.Config) (*App, error) {
	db, err := database.New(&cfg.DatabaseSecret)
	if err != nil {
		return nil, err
	}

	app := &App{fiber.New(), cfg, db}
	app.AttachMiddlewares()
	app.SetupRoutes()

	return app, nil
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

func (a *App) ApiV1(api fiber.Router) {
	v1, db := api.Group("/v1"), a.db
	agencyRepo, attendanceRepo, classRepo, registrationRepo, studentRepo, userRepo, roleRepo := agencies.NewRepo(
		db,
	), attendances.NewRepo(
		db,
	), classes.NewRepo(
		db,
	), registrations.NewRepo(
		db,
	), students.NewRepo(
		db,
	), users.NewRepo(
		db,
	), roles.NewRepo(
		db,
	)
	excel := spreadsheet.New()

	publicHandlers := []common.APIHandler{
		users.New(userRepo, a.config, a.config.ClientId, a.config.ClientSecret),
		registrations.New(registrationRepo),
		agencies.New(agencyRepo),
	}
	a.RegisterAPIHandlers(v1, publicHandlers)

	v1.Use(middlewares.ValidateJWT(a.config.Audience, a.config.Domain))

	privateHandlers := []common.APIHandler{
		agencies.New(agencyRepo),
		attendances.New(attendanceRepo, classRepo, excel, studentRepo, spreadsheet.NewExcelGenerator()),
		classes.New(classRepo),
		registrations.New(registrationRepo),
		students.New(studentRepo),
		users.New(userRepo, a.config, a.config.ClientId, a.config.Secret),
		roles.New(roleRepo, a.config, a.config.ClientId, a.config.Secret),
	}
	a.RegisterAPIHandlers(v1, privateHandlers)
}

func (a *App) AttachMiddlewares() {
	a.app.Use(recover.New())
	a.app.Use(mdwlogger.New())
	a.app.Use(cors.New())
}

func (a *App) SetupRoutes() {
	a.app.Get("/docs/*", swagger.HandlerDefault)
	a.app.Get("/healthz", healthz)
	a.app.Get("/", index)

	api := a.app.Group("/api")
	a.ApiV1(api)
}

func (a *App) App() *fiber.App {
	return a.app
}

func (a *App) Cleanup() {
	if err := a.db.Close(); err != nil {
		logger.Error("App.Cleanup err: ", err)
	}
}

func NewServerlessApp() (*App, error) {
	config, err := config.New()
	if err != nil {
		logger.Error("config.New err", "err", err)
		return nil, err
	}

	app, err := NewApp(config)
	if err != nil {
		logger.Error("NewApp err", "err", err)
		return nil, err
	}

	return app, nil
}
