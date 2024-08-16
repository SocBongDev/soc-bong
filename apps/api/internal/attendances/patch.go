package attendances

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/gofiber/fiber/v2"
)

// @PatchAttendance godoc
// @Summary Patch attendance api
// @Description Patch attendance
// @Tags Attendance
// @Accept json
// @Param post body []UpdateAttendanceRequest true "Patch attendance body"
// @Success 200 {object} Attendance
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances [patch]
func (h *AttendanceHandler) Patch(c *fiber.Ctx) error {
	body := []UpdateAttendanceRequest{}
	if err := c.BodyParser(&body); err != nil {
		log.Println("PatchAttendance.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("PatchAttendance request: %+v\n", body)
	req := make([]Attendance, len(body))
	for i, v := range body {
		req[i] = Attendance{
			CreateAttendanceRequest: CreateAttendanceRequest{AttendedStatus: v.AttendedStatus},
			BaseEntity:              common.BaseEntity{Id: v.Id},
		}
	}

	if err := h.repo.Update(req); err != nil {
		log.Println("PatchAttendance.Patch err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("PatchAttendance success. Response: %+v\n", body)

	return c.JSON(req)
}
