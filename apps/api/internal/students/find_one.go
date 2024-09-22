package students

import (
	"database/sql"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneStudent godoc
// @Summary Get student details api
// @Description Get one student
// @Tags Student
// @Accept json
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students/{id} [get]
func (h *StudentHandler) FindOne(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.InfoContext(ctx, "GetStudentDetails.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "GetStudentDetails id: ", id)
	resp := &Student{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(ctx, resp); err != nil {
		logger.ErrorContext(ctx, "GetStudentDetails.Query err", "err", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(resp)
}
