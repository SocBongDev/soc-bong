package attendances

import "github.com/gofiber/fiber/v2"

type AttendanceHandler struct {
	repo AttendanceRepository
}

func (h *AttendanceHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/attendances")
	r.Get("/", h.Find)
	r.Post("/", h.Insert)
	r.Put("/:id<int,min(1)>", h.Patch)
}

func New(repo AttendanceRepository) *AttendanceHandler {
	return &AttendanceHandler{repo}
}
