package attendances

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/entities"
)

type AttendanceRepository interface {
	Find(context.Context, *AttendanceQuery) ([]entities.Attendance, error)
	Insert(context.Context, []Attendance) error
	Update(context.Context, []Attendance) error
}
