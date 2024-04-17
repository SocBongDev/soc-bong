package students

import "github.com/gofiber/fiber/v2"

type StudentHandler struct {
	repo StudentRepository
}

func (h *StudentHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/students")
	r.Delete("/:id<int,min(1)>", h.Delete)
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Update)
}

func New(repo StudentRepository) *StudentHandler {
	return &StudentHandler{repo}
}
