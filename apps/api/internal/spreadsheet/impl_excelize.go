package spreadsheet

import (
	"bytes"
	"log"

	"github.com/SocBongDev/soc-bong/internal/attendances"
	"github.com/xuri/excelize/v2"
)

type spreadSheetExcelize struct{}

var _ SpreadSheet = (*spreadSheetExcelize)(nil)

func (s *spreadSheetExcelize) ExportClassAttendances(
	classAttendances []attendances.Attendance,
) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		log.Println("ExportExcel.NewSheet err: ", err)
		return nil, err
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	return f.WriteToBuffer()
}

func New() *spreadSheetExcelize {
	return &spreadSheetExcelize{}
}
