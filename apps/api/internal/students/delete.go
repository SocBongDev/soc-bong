package students

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @DeleteStudent godoc
// @Summary Delete student api
// @Description Delete student
// @Tags Student
// @Accept json
// @Param ids query []int true "Student IDs"
// @Success 200
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [delete]
func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	ctx, query := c.UserContext(), new(DeleteStudentQuery)
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "DeleteStudent.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "DeleteStudents request", "req", query)
	if err := h.repo.Delete(ctx, query.Ids); err != nil {
		logger.ErrorContext(ctx, "DeleteStudent.Delete err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(nil)
}
