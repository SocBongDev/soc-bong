package classes

import (
	"log"
	"strings"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateClass godoc
// @Summary Update class api
// @Description Update class
// @Tags Class
// @Accept json
// @Param post body WriteClassRequest true "Update class body"
// @Param id path int true "Class ID"
// @Success 200 {object} Class
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes/{id} [put]
func (h *ClassHandler) Update(c *fiber.Ctx) error {
	body := new(WriteClassRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateClass.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateClass.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateClass request: %+v\n", body)
	req := &Class{WriteClassRequest: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(req); err != nil {
		log.Println("UpdateClass.Update err: ", err)

		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateClass success. Response: %+v\n", body)

	return c.JSON(req)
}
