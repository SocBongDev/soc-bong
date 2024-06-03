package attendances

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @InsertAttendance godoc
// @Summary Create attendance api
// @Description Insert attendance
// @Tags Attendance
// @Accept json
// @Param post body []CreateAttendanceRequest true "Create attendance body"
// @Success 200
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances [post]
func (h *AttendanceHandler) Insert(c *fiber.Ctx) error {
	body := []CreateAttendanceRequest{}
	if err := c.BodyParser(&body); err != nil {
		log.Println("InsertAttendance.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("InsertAttendance request: %+v\n", body)
	req := make([]Attendance, len(body))
	for i, v := range body {
		req[i] = Attendance{CreateAttendanceRequest: v}
	}

	if err := h.repo.Insert(req); err != nil {
		log.Println("InsertAttendance.Insert err: ", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	log.Printf("InsertAttendance success. Response: %+v\n", body)

	return c.JSON(req)
}
