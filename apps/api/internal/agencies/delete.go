package agencies

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @DeleteAgency godoc
// @Summary Delete agency api
// @Description Delete agency
// @Tags Agency
// @Accept json
// @Param ids query []int true "Agency IDs"
// @Success 200
// @Failure 500 {string} string
// @Router /agencies [delete]
func (h *AgencyHandler) Delete(c *fiber.Ctx) error {
	query := new(DeleteAgencyQuery)
	if err := c.QueryParser(query); err != nil {
		log.Println("DeleteAgency.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("DeleteAgency request: %#v\n", query)

	if err := h.repo.Delete(query.Ids); err != nil {
		log.Println("DeleteAgency.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("DeleteAgency success. Response: %#v\n", query.Ids)

	return c.JSON(nil)
}
