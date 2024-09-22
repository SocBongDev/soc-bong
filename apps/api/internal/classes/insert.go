package classes

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertClass godoc
// @Summary Create class api
// @Description Insert class
// @Tags Class
// @Accept json
// @Param post body WriteClassRequest true "Create class body"
// @Success 200 {object} Class
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes [post]
func (h *ClassHandler) Insert(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteClassRequest)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "InsertClass.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "InsertClass request", "req", body)
	req := &Class{WriteClassRequest: *body}
	if err := h.repo.Insert(ctx, req); err != nil {
		logger.ErrorContext(ctx, "InsertClass.Insert err", "err", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
