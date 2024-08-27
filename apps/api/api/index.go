package handler

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SocBongDev/soc-bong/cmd/serve"
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/adaptor/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	config, err := config.New()
	if err != nil {
		logger.Error("config.New err", "err", err)
		panic(fmt.Sprintln("config.New err: ", err))
	}

	serverApp, err := serve.NewApp(config)
	if err != nil {
		logger.Error("NewApp err", "err", err)
		panic(fmt.Sprintln("NewApp err: ", err))
	}
	app := serverApp.App()

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", 5000)); err != nil {
			logger.Error("App.Listen err", "err", err)
			panic(fmt.Sprintln("App.Listen err: ", err))
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
	// redisConn.Close()
	logger.Info("Fiber was successful shutdown.")

	return adaptor.FiberApp(app)
}
