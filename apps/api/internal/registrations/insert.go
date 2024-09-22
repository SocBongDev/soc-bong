package registrations

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertRegistration godoc
// @Summary Create registration api
// @Description Insert registration
// @Tags Registration
// @Accept json
// @Param post body WriteRegistrationRequest true "Create registration body"
// @Success 200 {object} Registration
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations [post]
func (h *RegistrationHandler) Insert(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteRegistrationRequest)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "InsertRegistration.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "InsertRegistration request", "req", body)
	req := &Registration{WriteRegistrationRequest: *body}
	if err := h.repo.Insert(ctx, req); err != nil {
		logger.ErrorContext(ctx, "InsertRegistration.Insert err", "err", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
