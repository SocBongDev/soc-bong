package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindAgency godoc
// @Summary Get list agency details api
// @Description Get list agency
// @Tags Agency
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindAgencyResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /agencies [get]
func (h *AgencyHandler) Find(c *fiber.Ctx) error {
	query := &AgencyQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(c.Context(), "FindAgencies.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(c.Context(), "FindAgencies request", "req", query)
	data, err := h.repo.Find(c.Context(), query)
	if err != nil {
		logger.ErrorContext(c.Context(), "FindAgencies.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindAgencyResp{Data: data, Page: query.GetPage()}
	logger.DebugContext(c.Context(), "FindAgencies success", "resp", resp)

	return c.JSON(resp)
}
