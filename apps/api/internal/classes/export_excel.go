package classes

import (
	"fmt"
	"log"

	"github.com/SocBongDev/soc-bong/internal/attendances"
	"github.com/gofiber/fiber/v2"
)

// @FindOneClass godoc
// @Summary Get class excel api
// @Description Get one class excel file
// @Tags Class
// @Accept json
// @Param id path int true "Class ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /classes/{id}/export-excel [get]
func (h *ClassHandler) ExportExcel(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("ExportExcel.ParamsInt err: ", err)
		return fiber.ErrBadRequest
	}

	log.Println("ExportExcel id: ", id)

	classAttendances, err := h.attendanceRepo.Find(&attendances.AttendanceQuery{ClassId: id})
	if err != nil {
		log.Println("ExportExcel.Find err: ", err)
		return fiber.ErrInternalServerError
	}

	buf, err := h.spreadsheet.ExportClassAttendances(classAttendances)
	if err != nil {
		log.Println("f.WriteToBuffer err: ", err)
		return fiber.ErrInternalServerError
	}

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=Workbook-%d-.xlsx", id))
	c.Set("Content-Transfer-Encoding", "binary")
	c.Set("Expires", "0")
	return c.SendStream(buf)
}
