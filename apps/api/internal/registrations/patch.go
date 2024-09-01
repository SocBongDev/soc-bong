package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @PatchRegistration godoc
// @Summary Mark as done registration api
// @Description Mark a registration as processed
// @Tags Registration
// @Accept json
// @Param id path int true "Registration ID"
// @Success 200
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations/{id} [patch]
func (h *RegistrationHandler) Patch(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRegistration.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateRegistration request", "req", id)
	req := &Registration{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.MarkAsDone(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateRegistration.Update err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(nil)
}
