package users

import (
	"database/sql"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneUser godoc
// @Summary Get user details api
// @Description Get one user
// @Tags User
// @Accept json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h *UserHandler) FindOne(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "GetUserDetails.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "GetUserDetails request", "id", id)
	resp := &User{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(ctx, resp); err != nil {
		logger.ErrorContext(ctx, "GetUserDetails.Query err", "err", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(resp)
}
