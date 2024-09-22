package registrations

import (
	"database/sql"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneRegistration godoc
// @Summary Get registration details api
// @Description Get one registration
// @Tags Registration
// @Accept json
// @Param id path int true "Registration ID"
// @Success 200 {object} Registration
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations/{id} [get]
func (h *RegistrationHandler) FindOne(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "GetRegistrationDetails.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "GetRegistrationDetails request", "id", id)
	resp := &Registration{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(ctx, resp); err != nil {
		logger.ErrorContext(ctx, "GetRegistrationDetails.Query err", "err", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(resp)
}
