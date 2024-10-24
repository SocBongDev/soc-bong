package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RegistrationHandler struct {
	repo RegistrationRepository
}

func (h *RegistrationHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/registrations")
	r.Delete(
		"/",
		middlewares.ValidatePermissions([]string{"delete:registrations"}),
		h.Delete,
	)
	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:registrations"}),
		h.Find,
	)
	r.Get(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:registrations"}),
		h.FindOne,
	)
	r.Patch(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:registrations"}),
		h.Patch,
	)
	r.Post(
		"/",
		h.Insert,
	)
	r.Put(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:registrations"}),
		h.Update,
	)
}

func New(repo RegistrationRepository) common.APIHandler {
	return &RegistrationHandler{repo}
}
