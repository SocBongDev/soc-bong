package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AgencyHandler struct {
	repo AgencyRepository
}

func (h *AgencyHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/agencies")
	r.Delete("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"delete:agencies"}),
		h.Delete)
	r.Get("/",
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		h.Find,
	)
	r.Get("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		h.FindOne,
	)
	r.Patch("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:agencies"}),
		h.Patch,
	)
	r.Post("/",
		middlewares.ValidatePermissions([]string{"create:agencies"}),
		h.Insert,
	)
	r.Put("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:agencies"}),
		h.Update,
	)
}

func New(repo AgencyRepository) *AgencyHandler {
	return &AgencyHandler{repo}
}
