package students

import (
	"database/sql"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @FindOneStudent godoc
// @Summary Get student details api
// @Description Get one student
// @Tags Student
// @Accept json
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students/{id} [get]
func (h *StudentHandler) FindOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("GetStudentDetails.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("GetStudentDetails id: ", id)
	resp := &Student{BaseEntity: common.BaseEntity{Id: id}}
	if err := h.repo.FindOne(resp); err != nil {
		log.Println("GetStudentDetails.Query err: ", err)
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("GetStudentDetails success. Response: %+v\n", resp)

	return c.JSON(resp)
}
