package parents

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneParent godoc
// @Summary Get parent details api
// @Description Get one parent
// @Tags Parent
// @Accept json
// @Param id path int true "Parent ID"
// @Success 200 {object} Parent
// @Failure 500 {string} string
// @Router /parents/{id} [get]
func (h *ParentHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("GetParentDetails.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("GetParentDetails id: ", id)
	resp := &Parent{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("GetParentDetails.Query err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	log.Printf("GetParentDetails success. Response: %#v\n", resp)

	return c.JSON(resp)

}
