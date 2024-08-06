package users

import (
	"log"

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
	query := &UserQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindUser.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindUsers request: %+v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindUser.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindUserResp{Data: data, Page: query.GetPage(), PageSize: query.GetPageSize()}
	log.Printf("FindUsers success. Response: %+v\n", resp)
	return c.JSON(resp)
}
