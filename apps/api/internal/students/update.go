package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateStudent godoc
// @Summary Update student api
// @Description Update student
// @Tags Student
// @Accept json
// @Param post body WriteStudentRequest true "Update student body"
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students/{id} [put]
func (h *StudentHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteStudentRequest)
	if err := c.BodyParser(body); err != nil {
		logger.InfoContext(ctx, "UpdateStudent.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		logger.InfoContext(ctx, "UpdateStudent.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateStudent request", "req", body)
	req := &Student{
		WriteStudentRequest: *body,
		BaseEntity:          common.BaseEntity{Id: id},
	}
	if err := h.repo.Update(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateStudent.Update err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
