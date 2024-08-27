package handler

import (
	"log"
	"net/http"

	"github.com/SocBongDev/soc-bong/cmd/serve"
	"github.com/gofiber/adaptor/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	serverApp, err := serve.NewServerlessApp()
	if err != nil {
		log.Panic("NewApp err: ", err)
	}

	app := serverApp.App()
	return adaptor.FiberApp(app)
}
