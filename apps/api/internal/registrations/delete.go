package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @DeleteRegistration godoc
// @Summary Delete registration api
// @Description Delete registration
// @Tags Registration
// @Accept json
// @Param ids query []int true "Registration IDs"
// @Success 200
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations [delete]
func (h *RegistrationHandler) Delete(c *fiber.Ctx) error {
	ctx, query := c.UserContext(), new(DeleteRegistrationQuery)
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "DeleteRegistration.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "DeleteRegistrations request", "req", query)
	if err := h.repo.Delete(ctx, query.Ids); err != nil {
		logger.ErrorContext(ctx, "DeleteRegistration.Delete err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(nil)
}
