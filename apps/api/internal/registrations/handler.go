package registrations

import "github.com/gofiber/fiber/v2"

type RegistrationHandler struct {
	repo RegistrationRepository
}

func (h *RegistrationHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/registrations")
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
	r.Delete("/:id<int,min(1)>", h.Delete)
}

func New(repo RegistrationRepository) *RegistrationHandler {
	return &RegistrationHandler{repo}
}
