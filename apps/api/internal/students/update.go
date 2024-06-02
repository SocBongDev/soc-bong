package students

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @UpdateStudent godoc
// @Summary Update student api
// @Description Update student
// @Tags Student
// @Accept json
// @Param post body InsertStudentRequest true "Update student body"
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students/{id} [put]
func (h *StudentHandler) Update(c *fiber.Ctx) error {
	body := new(InsertStudentRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("UpdateStudent.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("UpdateStudent.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("UpdateStudent request: %+v\n", body)
	req := &Student{
		WriteStudentRequest: body.WriteStudentRequest,
		BaseEntity:          common.BaseEntity{Id: id},
	}
	if err := h.repo.Update(req); err != nil {
		log.Println("UpdateStudent.Update err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("UpdateStudent success. Response: %+v\n", body)

	return c.JSON(req)
}
