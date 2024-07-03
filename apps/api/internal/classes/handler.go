package classes

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	// attendanceRepo attendances.AttendanceRepository
	repo ClassRepository
	// spreadsheet    spreadsheet.SpreadSheet
}

func (h *ClassHandler) RegisterRoute(group fiber.Router) {
	r := group.Group("/classes")
	r.Get("/",
		middlewares.ValidatePermissions([]string{"read:classes"}),
		h.Find,
	)
	r.Get("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"read:classes"}),
		h.FindOne,
	)
	r.Post("/",
		middlewares.ValidatePermissions([]string{"create:classes"}),
		h.Insert,
	)
	r.Put("/:id<int,min(1)>",
		middlewares.ValidatePermissions([]string{"update:classes"}),
		h.Update,
	)
	r.Get("/:id<int,min(1)>/export-excel",
		h.ExportExcel,
	)
}

func New(
	// attendanceRepo attendances.AttendanceRepository,
	repo ClassRepository,
	// spreadsheet spreadsheet.SpreadSheet,
) common.APIHandler {
	// return &ClassHandler{attendanceRepo: attendanceRepo, repo: repo, spreadsheet: spreadsheet}
	return &ClassHandler{repo}
}
