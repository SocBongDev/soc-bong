package agencies

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @PatchAgency godoc
// @Summary Mark as done agency api
// @Description Mark a agency as processed
// @Tags Agency
// @Accept json
// @Param id path int true "Agency ID"
// @Success 200
// @Failure 500 {string} string
// @Router /agencies/{id} [patch]
func (h *AgencyHandler) Patch(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		log.Println("PatchAgency.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("PatchAgency id: %d\n", id)
	req := &Agency{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.MarkAsDone(req); err != nil {
		log.Println("PatchAgency.Patch err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("PatchAgency success. Response: %#v\n", req)

	return c.JSON(nil)
}
