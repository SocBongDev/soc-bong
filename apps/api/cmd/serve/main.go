package serve

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Args:  cobra.ArbitraryArgs,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Serving dinner...")

			config, err := config.New()
			if err != nil {
				logger.Error("config.New err", "err", err)
				panic(err)
			}

			serverApp, err := NewApp(config)
			if err != nil {
				logger.Error("NewApp err", "err", err)
				panic(err)
			}
			app := serverApp.app

			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", config.Port)); err != nil {
					logger.Error("app.Listen err", "err", err)
					panic(err)
				}
			}()

			c := make(chan os.Signal, 1)
			signal.Notify(
				c,
				os.Interrupt,
				syscall.SIGTERM,
			)

			_ = <-c
			logger.Info("Gracefully shutting down...")
			_ = app.Shutdown()

			logger.Info("Running cleanup tasks...")

			// Your cleanup tasks go here
			serverApp.db.Close()
			// redisConn.Close()
			logger.Info("Fiber was successful shutdown.")
		},
	}

	return cmd
}
