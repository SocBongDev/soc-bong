package students

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @DeleteStudent godoc
// @Summary Delete student api
// @Description Delete student
// @Tags Student
// @Accept json
// @Param ids query []int true "Student IDs"
// @Success 200
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [delete]
func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	query := new(DeleteStudentQuery)
	if err := c.QueryParser(query); err != nil {
		log.Println("DeleteStudent.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("DeleteStudents request: %#v\n", query)

	if err := h.repo.Delete(query.Ids); err != nil {
		log.Println("DeleteStudent.Delete err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("DeleteStudent success. Response: %#v\n", query.Ids)

	return c.JSON(nil)
}
