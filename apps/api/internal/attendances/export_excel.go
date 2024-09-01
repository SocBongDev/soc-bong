package attendances

import (
	"fmt"
	"strconv"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// @ExportExcel godoc
// @Summary Get class excel api
// @Description Get one class excel file
// @Tags Attendance
// @Accept json
// @Param id path int true "Class ID"
// @Param  period query string false "Time range"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /attendances/{id}/export-excel [get]
func (h *AttendanceHandler) ExportExcel(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id, err := c.ParamsInt("classId")
	if err != nil {
		logger.ErrorContext(ctx, "ExportExcel.ParamsInt err", "err", err)
		return fiber.ErrBadRequest
	}

	logger.InfoContext(ctx, "ExportExcel request", "id", id)
	query := &AttendanceQuery{}
	if err := c.QueryParser(query); err != nil {
		logger.ErrorContext(ctx, "FindAttendances.QueryParser err", "err", err)
		return fiber.ErrBadRequest
	}
	query.ClassId = id
	query.Format()

	attendanceResp, err := h.formatAttendances(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "ExportExcel.formatAttendances err", "err", err)
		return err
	}

	// buf, err := h.spreadsheet.ExportClassAttendances(attendanceResp)
	month, err := strconv.Atoi(query.month)
	if err != nil {
		logger.ErrorContext(ctx, "ExportExcel.Atoi err", "err", err)
		return err
	}

	year, err := strconv.Atoi(query.year)
	if err != nil {
		logger.ErrorContext(ctx, "ExportExcel.Atoi err", "err", err)
		return err
	}

	buf, err := h.excelGenerator.ExportClassAttendances(month, year, attendanceResp)
	if err != nil {
		logger.ErrorContext(ctx, "f.WriteToBuffer err", "err", err)
		return fiber.ErrInternalServerError
	}

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=Workbook-%d-.xlsx", id))
	c.Set("Content-Transfer-Encoding", "binary")
	c.Set("Expires", "0")
	return c.SendStream(buf)
}
