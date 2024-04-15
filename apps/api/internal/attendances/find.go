package attendances

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @FindAttendance godoc
// @Summary Get list attendance details api
// @Description Get list attendance
// @Tags Attendance
// @Accept json
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param  sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param  classId query string true "Class id"
// @Param  period query string false "Time range"
// @Success 200 {object} map[int][]Attendance
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances [get]
func (h *AttendanceHandler) Find(c *fiber.Ctx) error {
	query := &AttendanceQuery{}
	if err := c.QueryParser(query); err != nil {
		log.Println("FindAttendances.QueryParser err: ", err)
		return fiber.ErrBadRequest
	}
	query.Format()

	log.Printf("FindAttendances request: %#v\n", query)

	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindAttendances.All err: ", err)
		return fiber.ErrInternalServerError
	}

	resp := make(map[int][]Attendance)
	for _, a := range data {
		resp[a.StudentId] = append(resp[a.StudentId], a)
	}

	log.Printf("FindAttendances success. Response: %#v\n", resp)

	return c.JSON(resp)
}
