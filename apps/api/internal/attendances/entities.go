package attendances

import (
	"fmt"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
)

type CreateAttendanceRequest struct {
	AttendedAt common.DateTime `json:"attendedAt" swaggertype:"string"`
	ClassId    int             `json:"classId"`
	IsAttended bool            `json:"isAttended"`
	StudentId  int             `json:"studentId"`
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

func (q *AttendanceQuery) Format() {
	if q.Period == nil {
		now := time.Now()
		q.month, q.year = now.Format("01"), fmt.Sprint(now.Year())
		return
	}

	periodSplit := strings.Split(*q.Period, "-")
	q.month, q.year = periodSplit[0], periodSplit[1]
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
