package spreadsheet

import (
	"bytes"

	"github.com/SocBongDev/soc-bong/internal/entities"
)

type SpreadSheet interface {
	ExportClassAttendances(map[int]entities.AttendanceResponse) (*bytes.Buffer, error)
}
