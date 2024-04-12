package registrations

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @PatchRegistration godoc
// @Summary Mark as done registration api
// @Description Mark a registration as processed
// @Tags Registration
// @Accept json
// @Param id path int true "Registration ID"
// @Success 200
// @Failure 500 {string} string
// @Router /registrations/{id} [patch]
func (h *RegistrationHandler) Patch(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateRegistration.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateRegistration request: %#v\n", id)
	req := &Registration{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.MarkAsDone(req); err != nil {
		log.Println("UpdateRegistration.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateRegistration success. Response: %#v\n", id)

	return c.JSON(nil)
}
