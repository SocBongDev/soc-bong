package serve

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Args:  cobra.ArbitraryArgs,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Serving dinner...")

			app := fiber.New()
			config, err := config.New()
			if err != nil {
				log.Panicln("config.New err: ", err)
			}

			app.Get("/", func(ctx *fiber.Ctx) error {
				return ctx.JSON(fiber.Map{
					"uri":  ctx.Request().URI().String(),
					"path": ctx.Path(),
				})
			})

			app.Get(
				"/api/messages/admin",
				middlewares.ValidateJWT(config.Audience, config.Domain),
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
			// db.Close()
			// redisConn.Close()
			log.Println("Fiber was successful shutdown.")
		},
	}

	return cmd
}
