package users

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateUser godoc
// @Summary Update user api
// @Description Update user
// @Tags User
// @Accept json
// @Param post body UserInput true "Update user body"
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	ctx, body := c.UserContext(), new(UserInput)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateUser.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateUser.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateRegistration request: %+v\n", body)
	req := &User{UserInput: *body, BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Update(ctx, req); err != nil {
		log.Println("UpdateRegistration.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
