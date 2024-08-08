package spreadsheet

import (
	"bytes"
	"fmt"
	"time"

	"github.com/SocBongDev/soc-bong/internal/apperr"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/xuri/excelize/v2"
)

type ExcelGenerator struct {
	file *excelize.File
}

func NewExcelGenerator() *ExcelGenerator {
	return &ExcelGenerator{file: excelize.NewFile()}
}

func (e *ExcelGenerator) ExportClassAttendances(month, year int, classAttendances map[int]entities.AttendanceResponse) (*bytes.Buffer, error) {
	defer e.file.Close()

	if err := e.setupTemplate(); err != nil {
		logger.Error("ExportClassAttendances.setupTemplate err", "err", err)
		return nil, err
	}

	if err := e.writeStudentData(classAttendances); err != nil {
		logger.Error("ExportClassAttendances.writeStudentData err", "err", err)
		return e.file.WriteToBuffer()
	}

	if err := e.writeAttendanceData(classAttendances, month, year); err != nil {
		logger.Error("ExportClassAttendances.writeAttendanceData err", "err", err)
		return e.file.WriteToBuffer()
	}

	if err := e.writeFormulas(classAttendances); err != nil {
		logger.Error("ExportClassAttendances.writeFormulas err", "err", err)
		return e.file.WriteToBuffer()
	}

	return e.file.WriteToBuffer()
}

func (e *ExcelGenerator) setupTemplate() error {
	f, err := excelize.OpenFile("./internal/spreadsheet/template.xlsx")
	if err != nil {
		logger.Error("ExportClassAttendances.writeDataToExcel.OpenFile err", "err", err)
		return err
	}
	e.file = f

	return nil
}

func (e *ExcelGenerator) writeStudentData(classAttendances map[int]entities.AttendanceResponse) error {
	for idx, v := range classAttendances {
		rowIdx := idx + 3
		student := v.Student

		studentData := []struct {
			col   string
			value interface{}
		}{
			{"A", student.Class.Grade},
			{"B", student.LastName},
			{"C", student.FirstName},
			{"D", student.EnrolledAt.Time().Format(common.DayMonthYearLayout)},
			{"E", student.Dob.Time().Format(common.DayMonthYearLayout)},
			{"F", student.FatherPhoneNumber},
			{"G", student.MotherPhoneNumber},
			{"H", student.Zalo},
			{"I", student.Dob.Time().Year()},
			{"L", student.Ethnic},
			{"M", student.BirthPlace},
			{"N", student.FatherName},
			{"O", student.FatherDob.Time().Year()},
			{"P", student.FatherOccupation},
			{"Q", student.MotherName},
			{"R", student.MotherDob.Time().Year()},
			{"S", student.MotherOccupation},
			{"T", student.PermanentAddressCommune},
			{"U", student.PermanentAddressDistrict},
			{"V", student.PermanentAddressProvince},
			{"W", student.TempAddress},
			{"X", student.Landlord},
		}

		for _, data := range studentData {
			cell := fmt.Sprintf("%s%d", data.col, rowIdx)
			if err := e.file.SetCellValue(WORKSHEET, cell, data.value); err != nil {
				return apperr.New(fmt.Errorf("failed to set cell %s: %w", cell, err))
			}
		}

		// Set gender
		genderCol := map[bool]string{true: "J", false: "K"}[student.Gender]
		if err := e.file.SetCellValue(WORKSHEET, fmt.Sprintf("%s%d", genderCol, rowIdx), "x"); err != nil {
			return apperr.New(fmt.Errorf("failed to set gender: %w", err))
		}
	}
	return nil
}

func (e *ExcelGenerator) writeAttendanceData(classAttendances map[int]entities.AttendanceResponse, year, month int) error {
	// Generate all days in the specified month
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1)
	totalDays := endDate.Day()

	// Write attendance headers
	dateColIdx := 29 // Attendance data start here
	for day := 1; day <= totalDays; day++ {
		currentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		dayCell, _ := excelize.CoordinatesToCellName(dateColIdx+day-1, 1)
		weekdayCell, _ := excelize.CoordinatesToCellName(dateColIdx+day-1, 2)

		if err := e.file.SetCellValue(WORKSHEET, dayCell, day); err != nil {
			return apperr.New(fmt.Errorf("set day cell value error: %w", err))
		}

		vietnameseWeekday, ok := VietnameseWeekdayMap[currentDate.Weekday()]
		if !ok {
			vietnameseWeekday = "N/A"
		}

		if err := e.file.SetCellValue(WORKSHEET, weekdayCell, vietnameseWeekday); err != nil {
			return apperr.New(fmt.Errorf("set weekday cell value error: %w", err))
		}
	}

	// Write attendance data for each student
	rowIdx := 3
	for _, v := range classAttendances {
		// Create a map of attendance data for quick lookup
		attendanceMap := make(map[int]entities.AttendEnum)
		for _, att := range v.Attendances {
			attendanceMap[att.AttendedAt.Time().Day()] = att.AttendedStatus
		}

		for day := 1; day <= totalDays; day++ {
			cell, _ := excelize.CoordinatesToCellName(dateColIdx+day-1, rowIdx)

			if status, exists := attendanceMap[day]; exists {
				if err := e.file.SetCellValue(WORKSHEET, cell, int(status)); err != nil {
					return apperr.New(fmt.Errorf("set attendance status error: %w", err))
				}
			} else {
				// Leave cell empty for days without attendance data
				if err := e.file.SetCellValue(WORKSHEET, cell, ""); err != nil {
					return apperr.New(fmt.Errorf("set empty attendance cell error: %w", err))
				}
			}
		}
		rowIdx++
	}

	return nil
}

func (e *ExcelGenerator) writeFormulas(classAttendances map[int]entities.AttendanceResponse) error {
	rowCount := len(classAttendances)

	// First formula: Tăng ca
	colIdx := 29 + 31 // Assuming max 31 days in a month
	topLeftCell, _ := cellName(colIdx, 1)
	bottomRightCell, _ := cellName(colIdx+1, 1)

	if err := e.mergeCell(topLeftCell, bottomRightCell); err != nil {
		return apperr.New(fmt.Errorf("failed to merge cells for overtime: %w", err))
	}

	if err := e.file.SetCellValue(WORKSHEET, topLeftCell, fmt.Sprintf("Tăng ca T%02d/%d", time.Now().Month(), time.Now().Year())); err != nil {
		return apperr.New(fmt.Errorf("failed to set overtime cell value: %w", err))
	}

	overtimeStyle := &excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"ebb134"}},
	}
	if err := e.applyStyle(topLeftCell, overtimeStyle); err != nil {
		return apperr.New(fmt.Errorf("failed to apply overtime style: %w", err))
	}

	// Time and Dinner cells
	timeCell, _ := cellName(colIdx, 2)
	dinnerCell, _ := cellName(colIdx+1, 2)
	if err := e.file.SetCellValue(WORKSHEET, timeCell, "Giờ"); err != nil {
		return apperr.New(fmt.Errorf("failed to set time cell: %w", err))
	}
	if err := e.file.SetCellValue(WORKSHEET, dinnerCell, "Ăn tối"); err != nil {
		return apperr.New(fmt.Errorf("failed to set dinner cell: %w", err))
	}

	centerStyle := &excelize.Style{Alignment: &excelize.Alignment{Horizontal: "center"}}
	if err := e.applyStyle(timeCell, centerStyle); err != nil {
		return apperr.New(fmt.Errorf("failed to apply center style to time cell: %w", err))
	}
	if err := e.applyStyle(dinnerCell, centerStyle); err != nil {
		return apperr.New(fmt.Errorf("failed to apply center style to dinner cell: %w", err))
	}

	// Excused cell
	excuseColIdx := colIdx + 2
	excuseCell, _ := cellName(excuseColIdx, 2)
	if err := e.file.SetCellValue(WORKSHEET, excuseCell, "Phép"); err != nil {
		return apperr.New(fmt.Errorf("failed to set excuse cell: %w", err))
	}

	excuseStyle := &excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"FF0000"}},
	}
	if err := e.applyStyle(excuseCell, excuseStyle); err != nil {
		return apperr.New(fmt.Errorf("failed to apply excuse style: %w", err))
	}

	// Excused formula
	for i := 3; i < 3+rowCount; i++ {
		cell, _ := cellName(excuseColIdx, i)
		startCell, _ := cellName(29, i)
		endCell, _ := cellName(colIdx-1, i)
		formula := fmt.Sprintf(`COUNTIF(%s:%s, "P")`, startCell, endCell)
		if err := e.setCellFormula(cell, formula); err != nil {
			return apperr.New(fmt.Errorf("failed to set excuse formula: %w", err))
		}
	}

	// Second formula: CÁC KHOẢN PHẢI THU
	colIdx += 3
	topLeftCell, _ = cellName(colIdx, 1)
	bottomRightCell, _ = cellName(colIdx+9, 1)
	if err := e.mergeCell(topLeftCell, bottomRightCell); err != nil {
		return apperr.New(fmt.Errorf("failed to merge cells for payments: %w", err))
	}
	if err := e.file.SetCellValue(WORKSHEET, topLeftCell, fmt.Sprintf("CÁC KHOẢN PHẢI THU T%02d/%d", time.Now().Month(), time.Now().Year())); err != nil {
		return apperr.New(fmt.Errorf("failed to set payments cell value: %w", err))
	}

	paymentsStyle := &excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"FFA500"}},
	}
	if err := e.applyStyle(topLeftCell, paymentsStyle); err != nil {
		return apperr.New(fmt.Errorf("failed to apply payments style: %w", err))
	}

	// Subheaders
	subHeaders := []string{
		"NỢ T03", "TC T03", "CSVC", "ĐP", "Học toán",
		"Năng khiếu", "A.V", "Aerobic", "Tiền ăn T04", "HPT04",
	}
	subHeaderStyle := &excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"46a642"}},
	}
	for i, header := range subHeaders {
		cell, _ := cellName(colIdx+i, 2)
		if err := e.file.SetCellValue(WORKSHEET, cell, header); err != nil {
			return apperr.New(fmt.Errorf("failed to set subheader: %w", err))
		}
		if err := e.applyStyle(cell, subHeaderStyle); err != nil {
			return apperr.New(fmt.Errorf("failed to apply subheader style: %w", err))
		}
	}

	// Overtime calculation
	for i := 3; i < 3+rowCount; i++ {
		cell, _ := cellName(colIdx+1, i)
		hourCell, _ := cellName(colIdx-3, i)
		dinnerCell, _ := cellName(colIdx-2, i)
		formula := fmt.Sprintf(`%s*10+%s*10`, hourCell, dinnerCell)
		if err := e.setCellFormula(cell, formula); err != nil {
			return apperr.New(fmt.Errorf("failed to set overtime formula: %w", err))
		}
	}

	// Additional formulas (sum, collected, remaining) can be added here following the same pattern

	return nil
}

func (e *ExcelGenerator) applyStyle(cell string, style *excelize.Style) error {
	styleID, err := e.file.NewStyle(style)
	if err != nil {
		return apperr.New(fmt.Errorf("failed to create style: %w", err))
	}

	return e.file.SetCellStyle(WORKSHEET, cell, cell, styleID)
}

// Helper functions
func cellName(col, row int) (string, error) {
	return excelize.CoordinatesToCellName(col, row)
}

func (e *ExcelGenerator) setCellFormula(cell, formula string) error {
	return e.file.SetCellFormula(WORKSHEET, cell, formula)
}

func (e *ExcelGenerator) mergeCell(startCell, endCell string) error {
	return e.file.MergeCell(WORKSHEET, startCell, endCell)
}
