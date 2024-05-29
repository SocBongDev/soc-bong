package registrations

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneRegistration godoc
// @Summary Get registration details api
// @Description Get one registration
// @Tags Registration
// @Accept json
// @Param id path int true "Registration ID"
// @Success 200 {object} Registration
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /registrations/{id} [get]
func (h *RegistrationHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("GetRegistrationDetails.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("GetRegistrationDetails id: ", id)

	resp := &Registration{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("GetRegistrationDetails.Query err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("GetRegistrationDetails success. Response: %+v\n", resp)

	return c.JSON(resp)
}
