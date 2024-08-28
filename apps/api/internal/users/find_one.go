package users

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneUser godoc
// @Summary Get user details api
// @Description Get one user
// @Tags User
// @Accept json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h *UserHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("GetUserDetails.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("GetUserDetails id: ", id)

	resp := &User{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("GetUserDetails.Query err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("GetUserDetails success. Response: %+v\n", resp)

	return c.JSON(resp)
}
