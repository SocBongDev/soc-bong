package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateRegistration godoc
// @Summary Update registration api
// @Description Update registration
// @Tags Registration
// @Accept json
// @Param post body WriteRegistrationRequest true "Update registration body"
// @Param id path int true "Registration ID"
// @Success 200 {object} Registration
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations/{id} [put]
func (h *RegistrationHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteRegistrationRequest)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "UpdateRegistration.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "UpdateRegistration.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateRegistration request", "req", body)
	req := &Registration{WriteRegistrationRequest: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateRegistration.Update err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
