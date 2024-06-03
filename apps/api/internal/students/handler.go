package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/parents"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	repo       StudentRepository
	parentRepo parents.ParentRepository
}

func (h *StudentHandler) RegisterRoute(router fiber.Router) {
	r := router.Group("/students")
<<<<<<< HEAD

	r.Get(
		"/:id<int,min(1)>",
		h.FindOne,
	)

	r.Get(
		"/",
		h.Find,
	)
=======
	r.Get("/", h.Find)
	r.Get("/:id<int,min(1)>", h.FindOne)
	r.Post("/", h.Insert)
>>>>>>> d34fe274eb4f889099a7833b941bcaf0c65b76cd
}

func New(repo StudentRepository, parentRepo parents.ParentRepository) common.APIHandler {
	return &StudentHandler{repo: repo, parentRepo: parentRepo}
}
