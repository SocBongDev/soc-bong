package attendances

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/SocBongDev/soc-bong/internal/students"
	"github.com/gofiber/fiber/v2"
)

func (h *AttendanceHandler) formatAttendances(ctx context.Context, query *AttendanceQuery) (map[int]entities.AttendanceResponse, error) {
	data, err := h.repo.Find(ctx, query)
	if err != nil {
		logger.Error("FindAttendances.All err", "err", err)
		return nil, fiber.ErrInternalServerError
	}

	resp, studentIds := make(map[int]entities.AttendanceResponse), make([]int, 0)
	if len(data) == 0 {
		logger.Error("FindAttendances response is empty", "data", data)
		return resp, nil
	}

	for _, a := range data {
		attendanceResp, ok := resp[a.StudentId]
		if !ok {
			studentIds = append(studentIds, a.StudentId)
			resp[a.StudentId] = entities.AttendanceResponse{Attendances: []entities.Attendance{a}}
			continue
		}

		attendanceResp.Attendances = append(attendanceResp.Attendances, a)
		resp[a.StudentId] = attendanceResp
	}
	students, err := h.studentRepo.Find(ctx, &students.StudentQuery{Ids: studentIds})
	if err != nil {
		logger.Error("FindAttendances.studentRepo.Find err", "err", err)
		return nil, fiber.ErrInternalServerError
	}

	for _, student := range students {
		attendanceResp, ok := resp[student.Id]
		if !ok {
			logger.Error(
				"Something wrong. studentId doesn't exists in response map",
				"studentId", student.Id,
				"respMap", resp,
			)
			return nil, fiber.ErrInternalServerError
		}

		attendanceResp.Student = student
		resp[student.Id] = attendanceResp
	}

	return resp, nil
}
