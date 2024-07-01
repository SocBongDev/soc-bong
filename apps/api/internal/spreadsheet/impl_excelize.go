package spreadsheet

import (
	"bytes"
	"errors"
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

func CreateNumberCell(colIdx int, f *excelize.File, cellName, color string) int {
	no1StartCell, err := excelize.CoordinatesToCellName(colIdx, 1)
	if err != nil {
		log.Fatal("excelize.CoordinatesToCellName err: ", err)
	}

	no1EndCell, err := excelize.CoordinatesToCellName(colIdx+1, 2)
	if err != nil {
		log.Fatal("excelize.CoordinatesToCellName err: ", err)
	}

	f.MergeCell(WORKSHEET, no1StartCell, no1EndCell)
	f.SetCellStr(WORKSHEET, no1StartCell, cellName)
	numberCellStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "middle"},
		Border: []excelize.Border{
			{Type: "top", Style: 1, Color: "000000"},
			{Type: "bottom", Style: 1, Color: "000000"},
		},
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{color}},
	})
	if err != nil {
		log.Fatal(err)
	}
	f.SetCellStyle(WORKSHEET, no1StartCell, no1EndCell, numberCellStyle)
	f.SetColWidth(WORKSHEET, no1StartCell, no1StartCell, 15)
	colIdx += 2

	return colIdx
}

func (s *spreadSheetExcelize) ExportClassAttendances(
	classAttendances map[int]entities.AttendanceResponse,
) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Open template.xlsx
	f, err := excelize.OpenFile("./internal/spreadsheet/template.xlsx")
	if err != nil {
		log.Fatal("writeDataToExcel.OpenFile err: ", err)
	}

	// Set value of a cell
	rowIdx := 3
	isHandleFirstFormula := false
	for _, v := range classAttendances {
		f.SetCellValue(WORKSHEET, fmt.Sprintf("A%d", rowIdx), v.Student.Class.Grade)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("B%d", rowIdx), v.Student.LastName)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("C%d", rowIdx), v.Student.FirstName)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("D%d", rowIdx), v.Student.EnrolledAt.Time().Format(common.DayMonthYearLayout))
		f.SetCellValue(WORKSHEET, fmt.Sprintf("E%d", rowIdx), v.Student.Dob.Time().Format(common.DayMonthYearLayout))

		f.SetCellValue(WORKSHEET, fmt.Sprintf("F%d", rowIdx), v.Student.FatherPhoneNumber)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("G%d", rowIdx), v.Student.MotherPhoneNumber)

		f.SetCellValue(WORKSHEET, fmt.Sprintf("H%d", rowIdx), v.Student.Zalo)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("I%d", rowIdx), v.Student.Dob.Time().Year())

		if v.Student.Gender {
			f.SetCellValue(WORKSHEET, fmt.Sprintf("J%d", rowIdx), "x")
		} else {
			f.SetCellValue(WORKSHEET, fmt.Sprintf("K%d", rowIdx), "x")
		}

		f.SetCellValue(WORKSHEET, fmt.Sprintf("L%d", rowIdx), v.Student.Ethnic)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("M%d", rowIdx), v.Student.BirthPlace)

		f.SetCellValue(WORKSHEET, fmt.Sprintf("N%d", rowIdx), v.Student.FatherName)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("O%d", rowIdx), v.Student.FatherDob.Time().Year())
		f.SetCellValue(WORKSHEET, fmt.Sprintf("P%d", rowIdx), v.Student.FatherOccupation)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("Q%d", rowIdx), v.Student.MotherName)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("R%d", rowIdx), v.Student.MotherDob.Time().Year())
		f.SetCellValue(WORKSHEET, fmt.Sprintf("S%d", rowIdx), v.Student.MotherOccupation)

		f.SetCellValue(WORKSHEET, fmt.Sprintf("T%d", rowIdx), v.Student.PermanentAddressCommune)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("U%d", rowIdx), v.Student.PermanentAddressDistrict)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("V%d", rowIdx), v.Student.PermanentAddressProvince)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("W%d", rowIdx), v.Student.TempAddress)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("X%d", rowIdx), v.Student.Landlord)
		/* f.SetCellValue(WORKSHEET, fmt.Sprintf("Y%d", rowIdx),v.Student.)
		f.SetCellValue(WORKSHEET, fmt.Sprintf("Z%d", rowIdx), "N/A") */

		// Handle attendances data here
		dateColIdx := 29 // Attendance data start here
		for _, d := range v.Attendances {
			// Write day number to row 0 idx, date to row 1
			cell, err := excelize.CoordinatesToCellName(dateColIdx, 1)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}
			f.SetCellValue(WORKSHEET, cell, d.AttendedAt.Time().Day())

			cell, err = excelize.CoordinatesToCellName(dateColIdx, 2)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			vietnameseWeekday, ok := VietnameseWeekdayMap[d.AttendedAt.Time().Weekday().String()]
			if !ok {
				return nil, errors.New("Ain't noway this will happen lmao")
			}

			f.SetCellValue(
				WORKSHEET,
				cell,
				vietnameseWeekday,
			)
			dateColIdx++
		}

		colIdx := 29
		for _, d := range v.Attendances {
			// Fill in attendances status
			cell, err := excelize.CoordinatesToCellName(colIdx, rowIdx)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}
			f.SetCellValue(WORKSHEET, cell, int(d.AttendedStatus))
			colIdx++
		}

		// Formula time
		// First formular
		if !isHandleFirstFormula {
			topLeftCell, err := excelize.CoordinatesToCellName(colIdx, 1)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			bottomRightCell, err := excelize.CoordinatesToCellName(colIdx+1, 1)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
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
			}

			f.MergeCell(WORKSHEET, topLeftCell, bottomRightCell)
			f.SetCellValue(
				WORKSHEET,
				topLeftCell,
				fmt.Sprintf("Tăng ca T%02d/%d", 9, 2024),
			)
			f.SetCellStyle(WORKSHEET, topLeftCell, bottomRightCell, style)
			isHandleFirstFormula = true
		}

		horizontalCenterStyle, err := f.NewStyle(
			&excelize.Style{Alignment: &excelize.Alignment{Horizontal: "center"}},
		)
		if err != nil {
			log.Fatal("f.NewStyle err: ", err)
		}

		// Time cell
		cell, err := excelize.CoordinatesToCellName(colIdx, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}
		f.SetCellStr(WORKSHEET, cell, "Giờ")
		f.SetCellStyle(WORKSHEET, cell, cell, horizontalCenterStyle)

		// Dinner cell
		cell, err = excelize.CoordinatesToCellName(colIdx+1, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}
		f.SetCellStr(WORKSHEET, cell, "Ăn tối")
		f.SetCellStyle(WORKSHEET, cell, cell, horizontalCenterStyle)

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
			log.Fatal("f.NewStyle err: ", err)
		}

		excuseColIdx := colIdx + 2
		cell, err = excelize.CoordinatesToCellName(excuseColIdx, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}
		f.SetCellStr(WORKSHEET, cell, "Phép")
		f.SetCellStyle(WORKSHEET, cell, cell, pCellStyle)
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx+2, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			attendancesStartCell, err := excelize.CoordinatesToCellName(29, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			attendancesEndCell, err := excelize.CoordinatesToCellName(dateColIdx-1, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`COUNTIF(%s:%s; "P")`, attendancesStartCell, attendancesEndCell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
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
			log.Fatal("f.NewStyle err: ", err)
		}

		topLeftCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		bottomRightCell, err := excelize.CoordinatesToCellName(colIdx+9, 1)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		f.MergeCell(WORKSHEET, topLeftCell, bottomRightCell)
		f.SetCellStr(WORKSHEET, topLeftCell, fmt.Sprintf("CÁC KHOẢN PHẢI THU T%02d/%d", 9, 2023))
		f.SetCellStyle(WORKSHEET, topLeftCell, bottomRightCell, placeholderStyle)

		subHeaderStyle, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: "center"},
			Fill: excelize.Fill{
				Type: "pattern",
				// Pattern: 1,
				Color: []string{"46a642"},
			},
		})
		if err != nil {
			log.Fatal("f.NewStyle err: ", err)
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
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		subHeaderEndCell, err := excelize.CoordinatesToCellName(subHeaderEndIdx, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		for i, header := range subHeader {
			cell, err := excelize.CoordinatesToCellName(colIdx+i, 2)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			f.SetCellStr(WORKSHEET, cell, header)
		}
		f.SetCellStyle(WORKSHEET, subHeaderStartCell, subHeaderEndCell, subHeaderStyle)
		f.SetColWidth(WORKSHEET, subHeaderStartCell, subHeaderEndCell, 12)

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx+1, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			hourCell, err := excelize.CoordinatesToCellName(colIdx-3, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			dinnerCell, err := excelize.CoordinatesToCellName(colIdx-2, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`%s*10+%s*10`, hourCell, dinnerCell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
			}
		}
		colIdx += 10

		// Third formular
		minusEatMoneyIdx := colIdx
		minusEatMoneyCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		bottomRightCell, err = excelize.CoordinatesToCellName(colIdx+1, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		f.SetCellStr(WORKSHEET, minusEatMoneyCell, "Trừ tiền ăn")
		f.MergeCell(WORKSHEET, minusEatMoneyCell, bottomRightCell)
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
		}

		f.SetCellStyle(WORKSHEET, minusEatMoneyCell, bottomRightCell, style)
		f.SetColWidth(WORKSHEET, minusEatMoneyCell, bottomRightCell, 15)

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			nextCell, err := excelize.CoordinatesToCellName(colIdx+1, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			excuseCell, err := excelize.CoordinatesToCellName(excuseColIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			f.MergeCell(WORKSHEET, cell, nextCell)

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`+%s*30`, excuseCell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
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
			log.Fatal("f.NewStyle err: ", err)
		}

		sumCell, err := excelize.CoordinatesToCellName(colIdx, 1)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		collectCell, err := excelize.CoordinatesToCellName(colIdx, 2)
		if err != nil {
			log.Fatal("excelize.CoordinatesToCellName err: ", err)
		}

		f.SetCellStr(WORKSHEET, sumCell, "TỔNG")
		f.SetCellStr(WORKSHEET, collectCell, fmt.Sprintf("THU T%02d", 1))
		f.SetCellStyle(WORKSHEET, sumCell, collectCell, sumCellStyle)
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			subHeaderStartCell, err := excelize.CoordinatesToCellName(subHeaderStartIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			subHeaderEndCell, err := excelize.CoordinatesToCellName(subHeaderEndIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			minusEatMoneyCell, err := excelize.CoordinatesToCellName(minusEatMoneyIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`SUM(%s:%s)-%s`, subHeaderStartCell, subHeaderEndCell, minusEatMoneyCell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
			}
		}
		colIdx++

		// Fifth formular
		colIdx = CreateNumberCell(colIdx, f, "SỐ 2/03", "425da6")
		number1ColIdx := colIdx
		colIdx = CreateNumberCell(colIdx, f, "SỐ 1/04", "425da6")
		colIdx = CreateNumberCell(colIdx, f, "SỐ 2/04", "425da6")

		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			number1Cell, err := excelize.CoordinatesToCellName(number1ColIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			number3Cell, err := excelize.CoordinatesToCellName(colIdx-1, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`SUM(%s:%s)`, number1Cell, number3Cell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
			}
		}
		collectedIdx := colIdx
		colIdx = CreateNumberCell(colIdx, f, "ĐÃ THU", "a64258")

		// Sixth formular
		for i := 3; i < 3+len(classAttendances); i++ {
			cell, err := excelize.CoordinatesToCellName(colIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			sumCell, err := excelize.CoordinatesToCellName(sumCellIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			collectedCell, err := excelize.CoordinatesToCellName(collectedIdx, i)
			if err != nil {
				log.Fatal("excelize.CoordinatesToCellName err: ", err)
			}

			if err := f.SetCellFormula(
				WORKSHEET,
				cell,
				fmt.Sprintf(`%s-%s`, sumCell, collectedCell),
			); err != nil {
				log.Fatal("f.SetCellFormula err: ", err)
			}
		}
		colIdx = CreateNumberCell(colIdx, f, "CÒN NỢ", "ffa500")
		/* for i := 3; i < 3+len(resp.Data); i++ {

		   } */

		_ = CreateNumberCell(colIdx, f, "GHI CHÚ", "46a642")

		rowIdx++
	}

	// Save the spreadsheet
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Excel file created successfully!")

	return f.WriteToBuffer()
}

func New() *spreadSheetExcelize {
	return &spreadSheetExcelize{}
}
