package classes

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	repo ClassRepository
}

func (h *ClassHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/classes")
	r.Get("/", middlewares.ValidatePermissions([]string{"read:classes"}), h.Find)
	r.Get("/:id<int,min(1)>", middlewares.ValidatePermissions([]string{"read:classes"}), h.FindOne)
	r.Post("/", middlewares.ValidatePermissions([]string{"create:classes"}), h.Insert)
	r.Put("/:id<int,min(1)>", middlewares.ValidatePermissions([]string{"update:classes"}), h.Update)
}

func New(repo ClassRepository) common.APIHandler {
	return &ClassHandler{repo}
}
