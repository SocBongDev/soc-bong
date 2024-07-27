package agencies

import (
	"log"
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @InsertAgency godoc
// @Summary Create agency api
// @Description Insert agency
// @Tags Agency
// @Accept json
// @Param post body WriteAgencyRequest true "Create agency body"
// @Success 200 {object} Agency
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /agencies [post]
func (h *AgencyHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteAgencyRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertAgency.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(c.Context(), "InsertAgency request", "req", body)
	req := &Agency{WriteAgencyRequest: *body}
	if err := h.repo.Insert(req); err != nil {
		logger.ErrorContext(c.Context(), "InsertAgency.Insert err", "err", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	logger.DebugContext(c.Context(), "InsertAgency success", "resp", body)
	return c.JSON(req)
}
