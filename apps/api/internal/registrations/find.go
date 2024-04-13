package registrations

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type FindRegistrationResp common.FindResponse[Registration]

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
	query := &RegistrationQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindRegistrations.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindRegistrations request: %#v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindRegistrations.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindRegistrationResp{Data: data, Page: query.GetPage()}
	log.Printf("FindRegistrations success. Response: %#v\n", resp)

	return c.JSON(resp)
}
