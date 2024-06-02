package registrations

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @InsertRegistration godoc
// @Summary Create registration api
// @Description Insert registration
// @Tags Registration
// @Accept json
// @Param post body WriteRegistrationRequest true "Create registration body"
// @Success 200 {object} Registration
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations [post]
func (h *RegistrationHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteRegistrationRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertRegistration.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertRegistration request: %+v\n", body)
	req := &Registration{WriteRegistrationRequest: *body}
	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertRegistration.Insert err: ", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("InsertRegistration success. Response: %+v\n", body)

	return c.JSON(req)
}
