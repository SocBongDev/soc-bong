package classes

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateClass godoc
// @Summary Update class api
// @Description Update class
// @Tags Class
// @Accept json
// @Param post body WriteClassRequest true "Update class body"
// @Param id path int true "Class ID"
// @Success 200 {object} Class
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes/{id} [put]
func (h *ClassHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteClassRequest)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "UpdateClass.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "UpdateClass.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateClass request", "req", body)
	req := &Class{WriteClassRequest: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateClass.Update err", "err", err)

		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
