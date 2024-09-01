package attendances

import (
	"github.com/SocBongDev/soc-bong/internal/logger"
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
	query, ctx := &AttendanceQuery{}, c.UserContext()
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindAttendances.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}
	query.Format()

	logger.InfoContext(ctx, "FindAttendances request", "query", query)
	resp, err := h.formatAttendances(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindAttendances.formatAttendances err", "err", err)
		return err
	}

	return c.JSON(resp)
}
