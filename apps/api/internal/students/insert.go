package students

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertStudent godoc
// @Summary Create student api
// @Description Insert student
// @Tags Student
// @Accept json
// @Param post body WriteStudentRequest true "Create student body"
// @Success 200 {object} Student
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [post]
func (h *StudentHandler) Insert(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteStudentRequest)
	if err := c.BodyParser(body); err != nil {
		logger.InfoContext(ctx, "InsertStudent.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "InsertStudent request", "req", body)
	req := &Student{WriteStudentRequest: *body}
	if err := h.repo.Insert(ctx, req); err != nil {
		logger.InfoContext(ctx, "InsertStudent.Insert err", "err", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
