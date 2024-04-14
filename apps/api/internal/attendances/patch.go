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
// @Param post body UpdateAttendanceRequest true "Patch attendance body"
// @Param id path int true "Attendance ID"
// @Success 200 {object} Attendance
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances/{id} [patch]
func (h *AttendanceHandler) Patch(c *fiber.Ctx) error {
	body := new(UpdateAttendanceRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println("PatchAttendance.BodyParser err: ", err)
		return fiber.ErrBadRequest
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("PatchAttendance.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Printf("PatchAttendance request: %#v\n", body)
	req := &Attendance{
		CreateAttendanceRequest: CreateAttendanceRequest{IsAttended: body.IsAttended},
		BaseEntity:              common.BaseEntity{Id: id},
	}
	if err := h.repo.Update(req); err != nil {
		log.Println("PatchAttendance.Patch err: ", err)
		return fiber.ErrInternalServerError
	}

	log.Printf("PatchAttendance success. Response: %#v\n", body)

	return c.JSON(req)
}
