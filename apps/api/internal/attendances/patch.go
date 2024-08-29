package attendances

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
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
		logger.ErrorContext(c.UserContext(), "PatchAttendance.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(c.UserContext(), "PatchAttendance request", "req", body)
	req := make([]Attendance, len(body))
	for i, v := range body {
		req[i] = Attendance{
			CreateAttendanceRequest: CreateAttendanceRequest{AttendedStatus: v.AttendedStatus},
			BaseEntity:              common.BaseEntity{Id: v.Id},
		}
	}

	if err := h.repo.Update(req); err != nil {
		logger.ErrorContext(c.UserContext(), "PatchAttendance.Patch err", "err", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
