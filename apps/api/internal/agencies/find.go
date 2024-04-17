package agencies

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

type FindAgencyResponse common.FindResponse[Agency]

// @FindAgency godoc
// @Summary Get list agency details api
// @Description Get list agency
// @Tags Agency
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindAgencyResponse
// @Failure 500 {string} string
// @Router /agencies [get]
func (h *AgencyHandler) Find(c *fiber.Ctx) error {
	query := &AgencyQuery{}

	if err := c.QueryParser(query); err != nil {
		log.Println("FindAgencies.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindAgencies request: %#v\n", query)

	data, err := h.repo.Find(query)

	if err != nil {
		log.Println("FindAgencies.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindAgencyResponse{Data: data, Page: query.GetPage()}
	log.Printf("FindAgencies success. Response: %#v\n", resp)

	return c.JSON(resp)
}
