package agencies

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateAgency godoc
// @Summary Update acengy api
// @Description Update acengy
// @Tags Agency
// @Accept json
// @Param post body WriteAgencyRequest true "Update acengy body"
// @Param id path int true "Agency ID"
// @Success 200 {object} Agency
// @Failure 500 {string} string
// @Router /acengies/{id} [put]
func (h *AgencyHandler) Update(c *fiber.Ctx) error {
	body := new(WriteAgencyRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateAgency.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateAgency.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateAgency request: %#v\n", body)

	req := &Agency{BaseEntity: common.BaseEntity{Id: id}, WriteAgencyRequest: *body}
	if err := h.repo.Update(req); err != nil {
		log.Println("UpdateAgency.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateAgency success. Response: %#v\n", body)

	return c.JSON(req)
}
