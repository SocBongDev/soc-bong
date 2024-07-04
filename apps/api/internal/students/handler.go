package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	repo StudentRepository
}

func (h *StudentHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/students")
	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:students"}),
		h.Find,
	)
	r.Get(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:students"}),
		h.FindOne,
	)
	r.Post(
		"/",
		middlewares.ValidatePermissions([]string{"create:students"}),
		h.Insert,
	)
	r.Put(
		"/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:students"}),
		h.Update,
	)
}

func New(repo StudentRepository) common.APIHandler {
	return &StudentHandler{repo: repo}
}
