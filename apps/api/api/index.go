package handler

import (
	"fmt"
	"net/http"

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

	return adaptor.FiberApp(app)
}
