package classes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @FindClass godoc
// @Summary Get list class details api
// @Description Get list class
// @Tags Class
// @Accept json
// @Param  query query ClassQuery false "Query parameters"
// @Success 200 {object} FindClassResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes [get]
func (h *ClassHandler) Find(c *fiber.Ctx) error {
	query := &ClassQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindClasses.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindClasses request: %+v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindClasses.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindClassResp{Data: data, Page: query.GetPage()}
	log.Printf("FindClasses success. Response: %+v\n", resp)

	return c.JSON(resp)
}
