package spreadsheet

import (
	"bytes"
	"fmt"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/xuri/excelize/v2"
)

type spreadSheetExcelize struct{}

var (
	_                    SpreadSheet = (*spreadSheetExcelize)(nil)
	VietnameseWeekdayMap             = map[string]string{
		"Monday":    "T2",
		"Tuesday":   "T3",
		"Wednesday": "T4",
		"Thursday":  "T5",
		"Friday":    "T6",
		"Saturday":  "T7",
		"Sunday":    "CN",
	}
)

const WORKSHEET = "sheet1"

func createNumberCell(colIdx int, f *excelize.File, cellName, color string) (int, error) {
	no1StartCell, err := excelize.CoordinatesToCellName(colIdx, 1)
	if err != nil {
		log.Println("createNumberCell.excelize.CoordinatesToCellName err: ", err)
		return 0, err
	}

	no1EndCell, err := excelize.CoordinatesToCellName(colIdx+1, 2)
	if err != nil {
		log.Println("createNumberCell.excelize.CoordinatesToCellName err: ", err)
		return 0, err
	}

	if err := f.MergeCell(WORKSHEET, no1StartCell, no1EndCell); err != nil {
		log.Println("createNumberCell.MergeCell err: ", err)
		return 0, err
	}

	if err := f.SetCellStr(WORKSHEET, no1StartCell, cellName); err != nil {
		log.Println("createNumberCell.SetCellStr err: ", err)
		return 0, err
	}

	numberCellStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
		Border: []excelize.Border{
			{Type: "top", Style: 1, Color: "000000"},
			{Type: "bottom", Style: 1, Color: "000000"},
		},
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{color}},
	})
	if err != nil {
		log.Println("createNumberCell.NewStyle err: ", err)
		return 0, err
	}

	if err := f.SetCellStyle(WORKSHEET, no1StartCell, no1EndCell, numberCellStyle); err != nil {
		log.Println("createNumberCell.SetCellStyle err: ", err)
		return 0, err
	}

	if err := f.SetColWidth(WORKSHEET, no1StartCell, no1StartCell, 15); err != nil {
		log.Println("createNumberCell.SetColWidth err: ", err)
		// return 0, err
	}
	colIdx += 2

	return colIdx, nil
}

func (s *spreadSheetExcelize) ExportClassAttendances(
	classAttendances map[int]entities.AttendanceResponse,
) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("ExportClassAttendances.Close err: ", err)
		}
	}()

	// Open template.xlsx
	f, err := excelize.OpenFile("./internal/spreadsheet/template.xlsx")
	if err != nil {
		log.Println("ExportClassAttendances.writeDataToExcel.OpenFile err: ", err)
		return nil, err
	}

	// Set value of a cell
	rowIdx := 3
	isHandleFirstFormula := false
	for _, v := range classAttendances {
		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("A%d", rowIdx), v.Student.Class.Grade); err != nil {
			log.Println("ExportClassAttendances.SetStudentGrade err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("B%d", rowIdx), v.Student.LastName); err != nil {
			log.Println("ExportClassAttendances.SetStudentLastName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("C%d", rowIdx), v.Student.FirstName); err != nil {
			log.Println("ExportClassAttendances.SetStudentFirstName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("D%d", rowIdx), v.Student.EnrolledAt.Time().Format(common.DayMonthYearLayout)); err != nil {
			log.Println("ExportClassAttendances.SetStudentEnrolledAt err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("E%d", rowIdx), v.Student.Dob.Time().Format(common.DayMonthYearLayout)); err != nil {
			log.Println("ExportClassAttendances.SetStudentDob err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("F%d", rowIdx), v.Student.FatherPhoneNumber); err != nil {
			log.Println("ExportClassAttendances.SetFatherPhoneNumber err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("G%d", rowIdx), v.Student.MotherPhoneNumber); err != nil {
			log.Println("ExportClassAttendances.SetMortherPhoneNumber err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("H%d", rowIdx), v.Student.Zalo); err != nil {
			log.Println("ExportClassAttendances.SetParentZalo err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("I%d", rowIdx), v.Student.Dob.Time().Year()); err != nil {
			log.Println("ExportClassAttendances.SetStudentDobYear err: ", err)
			return f.WriteToBuffer()
		}

		genderCol := ""
		if v.Student.Gender {
			genderCol = fmt.Sprintf("J%d", rowIdx)
		} else {
			genderCol = fmt.Sprintf("K%d", rowIdx)
		}
		if err := f.SetCellValue(WORKSHEET, genderCol, "x"); err != nil {
			log.Println("ExportClassAttendances.SetStudentGender err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("L%d", rowIdx), v.Student.Ethnic); err != nil {
			log.Println("ExportClassAttendances.SetStudentEthnic err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("M%d", rowIdx), v.Student.BirthPlace); err != nil {
			log.Println("ExportClassAttendances.SetStudentBirthPlace err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("N%d", rowIdx), v.Student.FatherName); err != nil {
			log.Println("ExportClassAttendances.SetFatherName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("O%d", rowIdx), v.Student.FatherDob.Time().Year()); err != nil {
			log.Println("ExportClassAttendances.SetFatherDob err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("P%d", rowIdx), v.Student.FatherOccupation); err != nil {
			log.Println("ExportClassAttendances.SetFatherOccupation err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("Q%d", rowIdx), v.Student.MotherName); err != nil {
			log.Println("ExportClassAttendances.SetMotherName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("R%d", rowIdx), v.Student.MotherDob.Time().Year()); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("S%d", rowIdx), v.Student.MotherOccupation); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("T%d", rowIdx), v.Student.PermanentAddressCommune); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("U%d", rowIdx), v.Student.PermanentAddressDistrict); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("V%d", rowIdx), v.Student.PermanentAddressProvince); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("W%d", rowIdx), v.Student.TempAddress); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellValue(WORKSHEET, fmt.Sprintf("X%d", rowIdx), v.Student.Landlord); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}
		/* f.SetCellValue(WORKSHEET, fmt.Sprintf("Y%d", rowIdx),v.Student.)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("Z%d", rowIdx), "N/A") */

		// Handle attendances data here
		dateColIdx := 29 // Attendance data start here
		for _, d := range v.Attendances {
			// Write day number to row 0 idx, date to row 1
			cell, err := excelize.CoordinatesToCellName(dateColIdx, 1)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellValue(WORKSHEET, cell, d.AttendedAt.Time().Day()); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}

			cell, err = excelize.CoordinatesToCellName(dateColIdx, 2)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			vietnameseWeekday, ok := VietnameseWeekdayMap[d.AttendedAt.Time().Weekday().String()]
			if !ok {
				return f.WriteToBuffer()
			}

			if err := f.SetCellValue(WORKSHEET, cell, vietnameseWeekday); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}
			dateColIdx++
		}

		colIdx := 29
		for _, d := range v.Attendances {
			// Fill in attendances status
			cell, err := excelize.CoordinatesToCellName(colIdx, rowIdx)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellValue(WORKSHEET, cell, int(d.AttendedStatus)); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}
			colIdx++
		}

		// Formula time
		// First formular
		if !isHandleFirstFormula {
			topLeftCell, err := excelize.CoordinatesToCellName(colIdx, 1)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			bottomRightCell, err := excelize.CoordinatesToCellName(colIdx+1, 1)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			style, err := f.NewStyle(&excelize.Style{
				Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
				Fill: excelize.Fill{
					Type: "pattern",
					// Pattern: 1,
					Color: []string{"ebb134"},
				},
			})
			if err != nil {
				log.Println("f.NewStyle err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.MergeCell(WORKSHEET, topLeftCell, bottomRightCell); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellValue(
				WORKSHEET,
				topLeftCell,
				fmt.Sprintf("Tăng ca T%02d/%d", 9, 2024),
			); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellStyle(WORKSHEET, topLeftCell, bottomRightCell, style); err != nil {
				log.Println("ExportClassAttendances. err: ", err)
				return f.WriteToBuffer()
			}
			isHandleFirstFormula = true
		}

		horizontalCenterStyle, err := f.NewStyle(
			&excelize.Style{Alignment: &excelize.Alignment{Horizontal: "center"}},
		)
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		// Time cell
		cell, err := excelize.CoordinatesToCellName(colIdx, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, cell, "Giờ"); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, cell, cell, horizontalCenterStyle); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		// Dinner cell
		cell, err = excelize.CoordinatesToCellName(colIdx+1, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, cell, "Ăn tối"); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, cell, cell, horizontalCenterStyle); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		// P cell
		pCellStyle, err := f.NewStyle(
			&excelize.Style{
				Alignment: &excelize.Alignment{Horizontal: "center"},
				Fill: excelize.Fill{
					Type: "pattern",
					// Pattern: 1,
					Color: []string{"FF0000"},
				},
			},
		)
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		excuseColIdx := colIdx + 2
		cell, err = excelize.CoordinatesToCellName(excuseColIdx, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, cell, "Phép"); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, cell, cell, pCellStyle); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx+2, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			attendancesStartCell, err := excelize.CoordinatesToCellName(29, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			attendancesEndCell, err := excelize.CoordinatesToCellName(dateColIdx-1, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`COUNTIF(%s:%s; "P")`, attendancesStartCell, attendancesEndCell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		colIdx += 3

		// Second formular
		placeholderStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: "center"},
			Fill: excelize.Fill{
				Type: "pattern",
				// Pattern: 1,
				Color: []string{"FFA500"},
			},
		})
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		topLeftCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		bottomRightCell, err := excelize.CoordinatesToCellName(colIdx+9, 1)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.MergeCell(WORKSHEET, topLeftCell, bottomRightCell); err != nil {
			log.Println("ExportClassAttendances.Lmao err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, topLeftCell, fmt.Sprintf("CÁC KHOẢN PHẢI THU T%02d/%d", 9, 2023)); err != nil {
			log.Println("ExportClassAttendances.Kekw err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, topLeftCell, bottomRightCell, placeholderStyle); err != nil {
			log.Println("ExportClassAttendances.Kek err: ", err)
			return f.WriteToBuffer()
		}

		subHeaderStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: "center"},
			Fill: excelize.Fill{
				Type: "pattern",
				// Pattern: 1,
				Color: []string{"46a642"},
			},
		})
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		subHeader := []string{
			"NỢ T03",
			"TC T03",
			"CSVC",
			"ĐP",
			"Học toán",
			"Năng khiếu",
			"A.V",
			"Aerobic",
			"Tiền ăn T04",
			"HPT04",
		}
		subHeaderStartIdx, subHeaderEndIdx := colIdx, colIdx+len(subHeader)-1
		subHeaderStartCell, err := excelize.CoordinatesToCellName(subHeaderStartIdx, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		subHeaderEndCell, err := excelize.CoordinatesToCellName(subHeaderEndIdx, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		for i, header := range subHeader {
			cell, err := excelize.CoordinatesToCellName(colIdx+i, 2)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellStr(WORKSHEET, cell, header); err != nil {
				log.Println("ExportClassAttendances.SetCellStr err: ", err)
				return f.WriteToBuffer()
			}
		}

		if err := f.SetCellStyle(WORKSHEET, subHeaderStartCell, subHeaderEndCell, subHeaderStyle); err != nil {
			log.Println("ExportClassAttendances.SetCellStyle err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetColWidth(WORKSHEET, subHeaderStartCell, subHeaderEndCell, 12); err != nil {
			log.Println("ExportClassAttendances.Lord err: ", err)
			// return f.WriteToBuffer()
		}

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx+1, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			hourCell, err := excelize.CoordinatesToCellName(colIdx-3, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			dinnerCell, err := excelize.CoordinatesToCellName(colIdx-2, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`%s*10+%s*10`, hourCell, dinnerCell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		colIdx += 10

		// Third formular
		minusEatMoneyIdx := colIdx
		minusEatMoneyCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		bottomRightCell, err = excelize.CoordinatesToCellName(colIdx+1, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, minusEatMoneyCell, "Trừ tiền ăn"); err != nil {
			log.Println("ExportClassAttendances.SetEatMoneyStr err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.MergeCell(WORKSHEET, minusEatMoneyCell, bottomRightCell); err != nil {
			log.Println("ExportClassAttendances.MergeEatMoneyCell err: ", err)
			return f.WriteToBuffer()
		}
		style, err := f.NewStyle(
			&excelize.Style{
				Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
				// Border:    []excelize.Border{{}},
				Fill: excelize.Fill{
					Type: "pattern",
					// Pattern: 1,
					Color: []string{"FFA500"},
				},
			},
		)
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, minusEatMoneyCell, bottomRightCell, style); err != nil {
			log.Println("ExportClassAttendances.SetEatMoneyStyle err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetColWidth(WORKSHEET, minusEatMoneyCell, bottomRightCell, 15); err != nil {
			log.Println("ExportClassAttendances.SetEatMoneyColWidth err: ", err)
			// return f.WriteToBuffer()
		}

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			nextCell, err := excelize.CoordinatesToCellName(colIdx+1, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			excuseCell, err := excelize.CoordinatesToCellName(excuseColIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.MergeCell(WORKSHEET, cell, nextCell); err != nil {
				log.Println("ExportClassAttendances.Help me err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`+%s*30`, excuseCell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		colIdx += 2

		// Forth formular
		sumCellIdx := colIdx
		sumCellStyle, err := f.NewStyle(
			&excelize.Style{
				Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
				Font:      &excelize.Font{Bold: true},
				Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"8842a6"}},
			},
		)
		if err != nil {
			log.Println("f.NewStyle err: ", err)
			return f.WriteToBuffer()
		}

		sumCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		collectCell, err := excelize.CoordinatesToCellName(colIdx, 2)
		if err != nil {
			log.Println("excelize.CoordinatesToCellName err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, sumCell, "TỔNG"); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStr(WORKSHEET, collectCell, fmt.Sprintf("THU T%02d", 1)); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}

		if err := f.SetCellStyle(WORKSHEET, sumCell, collectCell, sumCellStyle); err != nil {
			log.Println("ExportClassAttendances. err: ", err)
			return f.WriteToBuffer()
		}
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			subHeaderStartCell, err := excelize.CoordinatesToCellName(subHeaderStartIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			subHeaderEndCell, err := excelize.CoordinatesToCellName(subHeaderEndIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			minusEatMoneyCell, err := excelize.CoordinatesToCellName(minusEatMoneyIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`SUM(%s:%s)-%s`, subHeaderStartCell, subHeaderEndCell, minusEatMoneyCell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		colIdx++

		// Fifth formular
		colIdx, err = createNumberCell(colIdx, f, "SỐ 2/03", "425da6")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}
		number1ColIdx := colIdx
		colIdx, err = createNumberCell(colIdx, f, "SỐ 1/04", "425da6")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}
		colIdx, err = createNumberCell(colIdx, f, "SỐ 2/04", "425da6")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			number1Cell, err := excelize.CoordinatesToCellName(number1ColIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			number3Cell, err := excelize.CoordinatesToCellName(colIdx-1, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`SUM(%s:%s)`, number1Cell, number3Cell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		collectedIdx := colIdx
		colIdx, err = createNumberCell(colIdx, f, "ĐÃ THU", "a64258")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}

		// Sixth formular
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			sumCell, err := excelize.CoordinatesToCellName(sumCellIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			collectedCell, err := excelize.CoordinatesToCellName(collectedIdx, i)
			if err != nil {
				log.Println("excelize.CoordinatesToCellName err: ", err)
				return f.WriteToBuffer()
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`%s-%s`, sumCell, collectedCell),
			); err != nil {
				log.Println("f.SetCellFormula err: ", err)
				return f.WriteToBuffer()
			}
		}
		colIdx, err = createNumberCell(colIdx, f, "CÒN NỢ", "ffa500")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}
		/* for i := 3; i < 3+len(resp.Data); i++ {

		   } */

		_, err = createNumberCell(colIdx, f, "GHI CHÚ", "46a642")
		if err != nil {
			log.Println("ExportClassAttendances.createNumberCell err: ", err)
			return f.WriteToBuffer()
		}

		rowIdx++
	}

	log.Println("Excel file created successfully!")
	return f.WriteToBuffer()
}

func New() *spreadSheetExcelize {
	return &spreadSheetExcelize{}
}
