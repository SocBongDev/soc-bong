package classes

import (
	"database/sql"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneClass godoc
// @Summary Get class details api
// @Description Get one class
// @Tags Class
// @Accept json
// @Param id path int true "Class ID"
// @Success 200 {object} Class
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes/{id} [get]
func (h *ClassHandler) FindOne(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "GetClassDetails.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "GetClassDetails id: ", id)
	resp := &Class{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(ctx, resp); err != nil {
		logger.ErrorContext(ctx, "GetClassDetails.Query err", "err", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(resp)
}
