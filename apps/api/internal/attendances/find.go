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

	log.Printf("FindAttendances request: %+v\n", query)
	resp, err := h.formatAttendances(query)
	if err != nil {
		log.Println("FindAttendances.formatAttendances err: ", err)
		return err
	}

	log.Printf("FindAttendances success. Response: %+v\n", resp)

	return c.JSON(resp)
}
