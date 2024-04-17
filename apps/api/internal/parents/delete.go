package parents

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @DeleteParent godoc
// @Summary Delete parent api
// @Description Delete parent
// @Tags Parent
// @Accept json
// @Param ids query []int true "Parent IDs"
// @Success 200
// @Failure 500 {string} string
// @Router /parents [delete]
func (h *ParentHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		log.Println("DeleteParent.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("DeleteParent id: ", id)
	req := &Parent{BaseEntity: common.BaseEntity{Id: id}}

	if err := h.repo.Delete(req); err != nil {
		log.Println("DeleteParent.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("DeleteParent success. Response: ", id)

	return c.JSON(req)
}
