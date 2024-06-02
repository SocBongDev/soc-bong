package students

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @FindStudent godoc
// @Summary Get list student details api
// @Description Get list student
// @Tags Student
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  search query string false "Search term"
// @Success 200 {object} FindStudentResp
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /students [get]
func (h *StudentHandler) Find(c *fiber.Ctx) error {
	query := &StudentQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindStudents.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("FindStudents request: %+v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindStudents.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := FindStudentResp{Data: data, Page: query.GetPage()}
	log.Printf("FindStudents success. Response: %+v\n", resp)

	return c.JSON(resp)
}
