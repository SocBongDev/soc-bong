package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindRegistration godoc
// @Summary Get list registration details api
// @Description Get list registration
// @Tags Registration
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindRegistrationResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations [get]
func (h *RegistrationHandler) Find(c *fiber.Ctx) error {
	ctx, query := c.UserContext(), &RegistrationQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindRegistrations.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "FindRegistrations request", "req", query)
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindRegistrations.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindRegistrationResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	return c.JSON(resp)
}
