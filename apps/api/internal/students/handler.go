package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	repo StudentRepository
}

func (h *StudentHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/students")
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
}

func New(repo StudentRepository) common.APIHandler {
	return &StudentHandler{repo: repo}
}
