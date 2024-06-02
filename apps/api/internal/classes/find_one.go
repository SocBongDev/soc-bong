package classes

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneClass godoc
// @Summary Get class details api
// @Description Get one class
// @Tags Class
// @Accept json
// @Param id path int true "Class ID"
// @Success 200 {object} Class
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes/{id} [get]
func (h *ClassHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("GetClassDetails.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("GetClassDetails id: ", id)
	resp := &Class{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("GetClassDetails.Query err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("GetClassDetails success. Response: %+v\n", resp)

	return c.JSON(resp)
}
