package attendances

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/SocBongDev/soc-bong/internal/students"
	"github.com/gofiber/fiber/v2"
)

func (h *AttendanceHandler) formatAttendances(query *AttendanceQuery) (map[int]entities.AttendanceResponse, error) {
	data, err := h.repo.Find(query)
	if err != nil {
		log.Println("FindAttendances.All err: ", err)
		return nil, fiber.ErrInternalServerError
	}

	resp, studentIds := make(map[int]entities.AttendanceResponse), make([]int, 0)
	if len(data) == 0 {
		log.Printf("FindAttendances response is empty. Data: %+v\n", data)
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
	students, err := h.studentRepo.Find(&students.StudentQuery{Ids: studentIds})
	if err != nil {
		log.Println("FindAttendances.studentRepo.Find err: ", err)
		return nil, fiber.ErrInternalServerError
	}

	for _, student := range students {
		attendanceResp, ok := resp[student.Id]
		if !ok {
			log.Printf("Something wrong. studentId: %d doesn't exists in response map: %+v\n", student.Id, resp)
			return nil, fiber.ErrInternalServerError
		}

		attendanceResp.Student = student
		resp[student.Id] = attendanceResp
	}

	return resp, nil
}
