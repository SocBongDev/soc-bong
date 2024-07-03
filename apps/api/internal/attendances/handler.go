package attendances

import (
	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/SocBongDev/soc-bong/internal/spreadsheet"
	"github.com/SocBongDev/soc-bong/internal/students"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	classRepo   classes.ClassRepository
	repo        AttendanceRepository
	spreadsheet spreadsheet.SpreadSheet
	studentRepo students.StudentRepository
}

func (h *AttendanceHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/attendances")
	r.Get(
		"/",
		middlewares.ValidatePermissions([]string{"read:attendances"}),
		h.Find,
	)
	r.Post(
		"/",
		middlewares.ValidatePermissions([]string{"create:attendances"}),
		h.Insert,
	)
	r.Patch(
		"/",
		middlewares.ValidatePermissions([]string{"update:attendances"}),
		h.Patch,
	)
	r.Get(
		"/:classId<int,min(1)>/export-excel",
		middlewares.ValidatePermissions([]string{"read:attendancess"}),
		h.ExportExcel,
	)
}

func New(repo AttendanceRepository, classRepo classes.ClassRepository, spreadsheet spreadsheet.SpreadSheet, studentRepo students.StudentRepository) common.APIHandler {
	return &AttendanceHandler{repo: repo, classRepo: classRepo, spreadsheet: spreadsheet, studentRepo: studentRepo}
}
