package attendances

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/SocBongDev/soc-bong/internal/students"
	"github.com/gofiber/fiber/v2"
)

type ClassAttendances struct {
	Data  map[int]entities.AttendanceResponse `json:"data"`
	Class *classes.Class                      `json:"class"`
}

func (h *AttendanceHandler) formatAttendances(ctx context.Context, query *AttendanceQuery) (*ClassAttendances, error) {
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.ErrorContext(ctx, "FindAttendances.All err", "err", err)
		return nil, fiber.ErrInternalServerError
	}

	classChan, classErrChan := make(chan *classes.Class, 1), make(chan error, 1)
	go func() {
		class := &classes.Class{BaseEntity: common.BaseEntity{Id: query.ClassId}}
		if err := h.classRepo.FindOne(ctx, class); err != nil {
			classErrChan <- err
			return
		}

		classChan <- class
	}()

	attendancesData, studentIds := make(map[int]entities.AttendanceResponse), make([]int, 0)
	// Process attendance data
	for _, a := range data {
		attendanceResp, ok := attendancesData[a.StudentId]
		if !ok {
			studentIds = append(studentIds, a.StudentId)
			attendancesData[a.StudentId] = entities.AttendanceResponse{Attendances: []entities.Attendance{a}}
			continue
		}
		attendanceResp.Attendances = append(attendanceResp.Attendances, a)
		attendancesData[a.StudentId] = attendanceResp
	}

	class := &classes.Class{}
	select {
	case class = <-classChan:
	case err := <-classErrChan:
		logger.ErrorContext(ctx, "FindAttendances.FindOne err", "err", err, "id", query.ClassId)
		return nil, fiber.ErrInternalServerError
	}

	resp := &ClassAttendances{Data: attendancesData, Class: class}
	if len(data) == 0 {
		logger.ErrorContext(ctx, "FindAttendances response is empty", "data", data)
		return resp, nil
	}

	// Find students concurrently
	studentsChan, errChan := make(chan []entities.Student), make(chan error)
	go func() {
		students, err := h.studentRepo.Find(ctx, &students.StudentQuery{Ids: studentIds})
		if err != nil {
			errChan <- err
			return
		}
		studentsChan <- students
	}()

	// Wait for student finding to complete
	select {
	case err := <-errChan:
		logger.ErrorContext(ctx, "FindAttendances.studentRepo.Find err", "err", err)
		return nil, err
	case students := <-studentsChan:
		for _, student := range students {
			attendanceResp, ok := attendancesData[student.Id]
			if !ok {
				logger.ErrorContext(
					ctx,
					"Something wrong. studentId doesn't exists in response map",
					"studentId", student.Id,
					"respMap", resp,
				)
				return nil, fiber.ErrInternalServerError
			}
			attendanceResp.Student = student
			attendancesData[student.Id] = attendanceResp
		}
	}

	return resp, nil
}
