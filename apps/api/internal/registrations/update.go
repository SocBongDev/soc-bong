package registrations

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateRegistration godoc
// @Summary Update registration api
// @Description Update registration
// @Tags Registration
// @Accept json
// @Param post body WriteRegistrationRequest true "Update registration body"
// @Param id path int true "Registration ID"
// @Success 200 {object} Registration
// @Failure 500 {string} string
// @Router /registrations/{id} [put]
func (h *RegistrationHandler) Update(c *fiber.Ctx) error {
	body := new(WriteRegistrationRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateRegistration.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateRegistration.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateRegistration request: %#v\n", body)
	req := &Registration{WriteRegistrationRequest: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(req); err != nil {
		log.Println("UpdateRegistration.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateRegistration success. Response: %#v\n", body)

	return c.JSON(req)
}
