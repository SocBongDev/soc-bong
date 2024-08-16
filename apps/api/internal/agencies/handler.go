package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AgencyHandler struct {
	repo AgencyRepository
}

func (h *AgencyHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/agencies")
	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		h.Find,
	)
	r.Get(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:agencies"}),
		h.FindOne,
	)
	r.Post(
		"/",
		middlewares.ValidatePermissions([]string{"create:agencies"}),
		h.Insert,
	)
	r.Put(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:agencies"}),
		h.Update,
	)
}

func New(repo AgencyRepository) common.APIHandler {
	return &AgencyHandler{repo}
}
