package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	repo StudentRepository
}

func (h *StudentHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/students")

	r.Get(
		"/:id<int,min(1)>",
		h.FindOne,
	)

	r.Get(
		"/",
		h.Find,
	)
}

func New(repo StudentRepository) common.APIHandler {
	return &StudentHandler{repo}
}
