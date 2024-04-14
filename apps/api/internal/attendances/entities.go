package attendances

import (
	"fmt"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
)

type CreateAttendanceRequest struct {
	AttendedAt time.Time `json:"attendedAt"`
	ClassId    int       `json:"classId"`
	IsAttended bool      `json:"isAttended"`
	StudentId  int       `json:"studentId"`
}

type UpdateAttendanceRequest struct {
	Id         int  `json:"id"`
	IsAttended bool `json:"isAttended"`
}

type AttendanceQuery struct {
	ClassId int     `json:"classId"`
	Period  *string `json:"period"`
	month   string  `json:"-"`
	year    string  `json:"-"`
	common.Sorter
}

func NewAttendanceQuery(classId int, period *string) *AttendanceQuery {
	if period == nil {
		now := time.Now()
		return &AttendanceQuery{
			ClassId: classId,
			month:   now.Month().String(),
			year:    fmt.Sprint(now.Year()),
		}
	}

	periodSplit := strings.Split(*period, "-")
	return &AttendanceQuery{ClassId: classId, month: periodSplit[0], year: periodSplit[1]}
}

func (q AttendanceQuery) PeriodMonth() string {
	return q.month
}

func (q AttendanceQuery) PeriodYear() string {
	return q.year
}

type Attendance struct {
	common.BaseEntity
	CreateAttendanceRequest
}

func (e *Attendance) TableName() string {
	return "attendances"
}
