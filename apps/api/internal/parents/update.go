package parents

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateParent godoc
// @Summary Update parent api
// @Description Update parent
// @Tags Parent
// @Accept json
// @Param post body WriteParentRequest true "Update parent body"
// @Param id path int true "Parent ID"
// @Success 200 {object} Parent
// @Failure 500 {string} string
// @Router /parents/{id} [put]
func (h *ParentHandler) Update(c *fiber.Ctx) error {
	body := new(WriteParentRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateParent.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateParent.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateParent request: %#v\n", body)

	req := &Parent{BaseEntity: common.BaseEntity{Id: id}, WriteParentRequest: *body}

	if err := h.repo.Update(req); err != nil {
		log.Println("UpdateParent.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateParent success. Response: %#v\n", body)

	return c.JSON(req)
}
