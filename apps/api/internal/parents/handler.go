package parents

import "github.com/gofiber/fiber/v2"

type ParentHandler struct {
	repo ParentRepository
}

func (h *ParentHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/parents")
	r.Delete("/:id<int,min(1)>", h.Delete)
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
}

func New(repo ParentRepository) *ParentHandler {
	return &ParentHandler{repo}
}
