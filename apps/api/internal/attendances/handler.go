package attendances

import (
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	repo AttendanceRepository
}

func (h *AttendanceHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/attendances")
	r.Get("/", middlewares.ValidatePermissions([]string{"read:attendances"}), h.Find)
	r.Post("/", middlewares.ValidatePermissions([]string{"create:attendances"}), h.Insert)
	r.Put(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:attendances"}),
		h.Patch,
	)
}

func New(repo AttendanceRepository) *AttendanceHandler {
	return &AttendanceHandler{repo}
}
