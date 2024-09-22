package users

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateUser godoc
// @Summary Update user api
// @Description Update user
// @Tags User
// @Accept json
// @Param post body UserInput true "Update user body"
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(UserInput)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "UpdateUser.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "UpdateUser.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateUser.Request Body request", "req", body)
	req := &User{UserInput: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateUser.Update err", "err", err)
		return fiber.ErrInternalServerError
	}
	logger.InfoContext(ctx, "UpdateUser.Update Success Response", "req", req)
	return c.JSON(req)
}
