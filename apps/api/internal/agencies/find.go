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
	ctx := c.UserContext()
	query := &AgencyQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindAgencies.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "FindAgencies request", "req", query)
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindAgencies.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindAgencyResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	return c.JSON(resp)
}
