package users

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @FindUser godoc
// @Summary Get list user details api
// @Description Get list user
// @Tags User
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Param  email query string false "Email term"
// @Success 200 {object} FindUserResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users [get]
func (h *UserHandler) Find(c *fiber.Ctx) error {
	ctx, query := c.UserContext(), &UserQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindUsers.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "FindUsers request", "req", query)
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindUser.All err", "err", err)
		return fiber.ErrInternalServerError
	}

	resp := FindUserResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	logger.InfoContext(ctx, "FindUsers.Success Response", resp)
	return c.JSON(resp)
}
