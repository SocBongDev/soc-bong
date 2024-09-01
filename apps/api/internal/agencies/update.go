package agencies

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @UpdateAgency godoc
// @Summary Update agency api
// @Description Update agency
// @Tags Agency
// @Accept json
// @Param post body WriteAgencyRequest true "Update agency body"
// @Param id path int true "Agency ID"
// @Success 200 {object} Agency
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /agencies/{id} [put]
func (h *AgencyHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(WriteAgencyRequest)
	if err := c.BodyParser(body); err != nil {
		logger.ErrorContext(ctx, "UpdateAgency.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(ctx, "UpdateAgency.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "UpdateAgency request", "req", body)
	req := &Agency{WriteAgencyRequest: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(ctx, req); err != nil {
		logger.ErrorContext(ctx, "UpdateAgency.Update err", "err", err)

		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
