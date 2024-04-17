package classes

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	repo ClassRepository
}

func (h *ClassHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/classes")
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
}

func New(repo ClassRepository) common.APIHandler {
	return &ClassHandler{repo}
}
