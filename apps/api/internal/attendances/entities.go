package attendances

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
)

type AttendEnum int

const (
	Absented AttendEnum = iota
	Attended
	Excused
	Dayoff
	Holiday
)

var (
	_ json.Marshaler   = (*AttendEnum)(nil)
	_ json.Unmarshaler = (*AttendEnum)(nil)
)

func (e AttendEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *AttendEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*e, err = parseAttendEnum(s)
	return err
}

func (e AttendEnum) String() string {
	switch e {
	case Absented:
		return "absented"
	case Attended:
		return "attended"
	case Excused:
		return "excused"
	case Dayoff:
		return "dayoff"
	case Holiday:
		return "holiday"
	default:
		return "unknown"
	}
}

func parseAttendEnum(s string) (AttendEnum, error) {
	switch s {
	case "absented":
		return Absented, nil
	case "attended":
		return Attended, nil
	case "excused":
		return Excused, nil
	case "dayoff":
		return Dayoff, nil
	case "holiday":
		return Holiday, nil
	default:
		return 0, errors.New("invalid AttendEnum value")
	}
}

type CreateAttendanceRequest struct {
	AttendedAt     common.DateTime `json:"attendedAt"     swaggertype:"string"`
	AttendedStatus AttendEnum      `json:"attendedStatus" swaggertype:"string" enums:"absented,attended,excused,dayoff,holiday"`
	ClassId        int             `json:"classId"`
	StudentId      int             `json:"studentId"`
}

type UpdateAttendanceRequest struct {
	Id             int        `json:"id"`
	AttendedStatus AttendEnum `json:"attendedStatus"`
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
