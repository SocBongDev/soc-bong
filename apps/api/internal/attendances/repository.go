package attendances

import "github.com/SocBongDev/soc-bong/internal/entities"

type AttendanceRepository interface {
	Find(*AttendanceQuery) ([]entities.Attendance, error)
	Insert([]Attendance) error
	Update([]Attendance) error
}
