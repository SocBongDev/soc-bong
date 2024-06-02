package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type AgencyHandler struct {
	repo AgencyRepository
}

func (h *AgencyHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/agencies")
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
}

func New(repo AgencyRepository) common.APIHandler {
	return &AgencyHandler{repo}
}
