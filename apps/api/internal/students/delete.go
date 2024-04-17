package students

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @DeleteStudent godoc
// @Summary Delete student api
// @Description Delete student
// @Tags Student
// @Accept json
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students/{id} [delete]
func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("DeleteStudent.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	req := &Student{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.Delete(req); err != nil {
		log.Println("DeleteStudent.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("DeleteStudent success. Response: %#v\n", req)

	return c.JSON(req)
}
