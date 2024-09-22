package spreadsheet

import "github.com/xuri/excelize/v2"

func (s *spreadSheetExcelize) generateSheet(f *excelize.File) {
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
}
