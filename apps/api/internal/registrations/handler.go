package registrations

import "github.com/gofiber/fiber/v2"

type RegistrationHandler struct {
	repo RegistrationRepository
}

func (h *RegistrationHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/registrations")
	r.Get("/", h.Find)
}

func New(repo RegistrationRepository) *RegistrationHandler {
	return &RegistrationHandler{repo}
}
