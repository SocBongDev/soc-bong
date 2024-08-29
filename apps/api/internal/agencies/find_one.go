package agencies

import (
	"database/sql"

	"github.com/SocBongDev/soc-bong/internal/apperr"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindOneAgency godoc
// @Summary Get agency details api
// @Description Get one agency
// @Tags Agency
// @Accept json
// @Param id path int true "Agency ID"
// @Success 200 {object} Agency
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /agencies/{id} [get]
func (h *AgencyHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.ErrorContext(c.Context(), "GetAgencyDetails.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(c.Context(), "GetAgencyDetails request", " id", id)
	resp := &Agency{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(c.Context(), resp); err != nil {
		logger.ErrorContext(c.Context(), "GetAgencyDetails.Query err", "err", apperr.New(err))
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	logger.DebugContext(c.Context(), "GetAgencyDetails success", "resp", resp)
	return c.JSON(resp)
}
