package parents

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @InsertParent godoc
// @Summary Create parent api
// @Description Insert parent
// @Tags Parent
// @Accept json
// @Param post body WriteParentRequest true "Create parent body"
// @Success 200 {object} Parent
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Router /parents [post]
func (h *ParentHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteParentRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	log.Printf("InsertParent request: %#v\n", body)
	req := &Parent{WriteParentRequest: *body}
	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertParent.Insert err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("InsertParent success. Response: %#v\n", body)

	return c.JSON(req)
}
