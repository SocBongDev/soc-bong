package classes

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @InsertClass godoc
// @Summary Create class api
// @Description Insert class
// @Tags Class
// @Accept json
// @Param post body WriteClassRequest true "Create class body"
// @Success 200 {object} Class
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes [post]
func (h *ClassHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteClassRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertClass.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertClass request: %#v\n", body)
	req := &Class{WriteClassRequest: *body}
	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertClass.Insert err: ", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("InsertClass success. Response: %#v\n", body)

	return c.JSON(req)
}
