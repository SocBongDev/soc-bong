package registrations

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @DeleteRegistration godoc
// @Summary Delete registration api
// @Description Delete registration
// @Tags Registration
// @Accept json
// @Param ids query []int true "Registration IDs"
// @Success 200
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations [delete]
func (h *RegistrationHandler) Delete(c *fiber.Ctx) error {
	query := new(DeleteRegistrationQuery)
	if err := c.QueryParser(query); err != nil {
		log.Println("DeleteRegistration.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("DeleteRegistrations request: %+v\n", query)

	if err := h.repo.Delete(query.Ids); err != nil {
		log.Println("DeleteRegistration.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("DeleteRegistration success. Response: %+v\n", query.Ids)

	return c.JSON(nil)
}
