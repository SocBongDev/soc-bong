package attendances

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
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
	ctx, body := c.UserContext(), []CreateAttendanceRequest{}
	if err := c.BodyParser(&body); err != nil {
		logger.ErrorContext(ctx, "InsertAttendance.BodyParser err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "InsertAttendance request", "request", body)
	req := make([]Attendance, len(body))
	for i, v := range body {
		req[i] = Attendance{CreateAttendanceRequest: v}
	}

	if err := h.repo.Insert(ctx, req); err != nil {
		logger.ErrorContext(ctx, "InsertAttendance.Insert err", "err", err)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fiber.ErrUnprocessableEntity
		}

		return fiber.ErrInternalServerError
	}

	return c.JSON(req)
}
