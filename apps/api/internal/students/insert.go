package students

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @InsertStudent godoc
// @Summary Create student api
// @Description Insert student
// @Tags Student
// @Accept json
// @Param post body WriteStudentRequest true "Create student body"
// @Success 200 {object} Student
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [post]
func (h *StudentHandler) Insert(c *fiber.Ctx) error {
	body := new(WriteStudentRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("InsertStudent.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertStudent request: %+v\n", body)
	req := &Student{WriteStudentRequest: *body}

	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertStudent.Insert err: ", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("InsertStudent success. Response: %+v\n", body)

	return c.JSON(req)
}
