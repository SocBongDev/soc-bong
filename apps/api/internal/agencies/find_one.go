package agencies

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneAgency godoc
// @Summary Get agency details api
// @Description Get one agency
// @Tags Agency
// @Accept json
// @Param id path int true "Agency ID"
// @Success 200 {object} Agency
// @Failure 500 {string} string
// @Router /agencies/{id} [get]
func (h *AgencyHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		log.Println(("FindOneAgency.ParamsInt err: "), err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindOneAgency id: %d\n", id)

	resp := &Agency{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("FindOneAgency.Query err: ", err)

		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	log.Printf("FindOneAgency success. Response: %#v\n", resp)

	return c.JSON(resp)
}
