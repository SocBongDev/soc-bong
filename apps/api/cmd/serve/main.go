package serve

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Args:  cobra.ArbitraryArgs,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Serving dinner...")

			config, err := config.New()
			if err != nil {
				log.Panicln("config.New err: ", err)
			}

			serverApp, err := NewApp(config)
			if err != nil {
				log.Panic("NewApp err: ", err)
			}
			app := serverApp.app

			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", 5000)); err != nil {
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
			serverApp.db.Close()
			// redisConn.Close()
			log.Println("Fiber was successful shutdown.")
		},
	}

	return cmd
}
