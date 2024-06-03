package agencies

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @InsertAgency godoc
// @Summary Create agency api
// @Description Insert agency
// @Tags Agency
// @Accept json
// @Param post body WriteAgencyRequest true "Create agency body"
// @Success 200 {object} Agency
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /agencies [post]
func (h *AgencyHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteAgencyRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertAgency.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertAgency request: %+v\n", body)
	req := &Agency{WriteAgencyRequest: *body}
	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertAgency.Insert err: ", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("InsertAgency success. Response: %+v\n", body)

	return c.JSON(req)
}
