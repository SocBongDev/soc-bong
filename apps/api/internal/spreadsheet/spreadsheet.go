package spreadsheet

import (
	"bytes"

	"github.com/SocBongDev/soc-bong/internal/attendances"
)

type SpreadSheet interface {
	ExportClassAttendances(classAttendances []attendances.Attendance) (*bytes.Buffer, error)
}
