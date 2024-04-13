package registrations

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @DeleteRegistration godoc
// @Summary Delete registration api
// @Description Delete registration
// @Tags Registration
// @Accept json
// @Param id path int true "Registration ID"
// @Success 200 {object} Registration
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations/{id} [delete]
func (h *RegistrationHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("DeleteRegistration.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	req := &Registration{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Delete(req); err != nil {
		log.Println("DeleteRegistration.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("DeleteRegistration success. Response: %#v\n", req)

	return c.JSON(req)
}
