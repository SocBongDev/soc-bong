package entities

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

type DbAttendanceResult struct {
	common.BaseEntity
	CreateAttendanceRequest

	BirthPlace               string          `db:"birth_place"`
	StudentCreatedAt         time.Time       `db:"student_created_at"`
	StudentUpdatedAt         time.Time       `db:"student_updated_at"`
	Dob                      common.DateTime `db:"dob"`
	EnrolledAt               common.DateTime `db:"enrolled_at"`
	Ethnic                   string          `db:"ethnic"`
	FirstName                string          `db:"first_name"`
	Gender                   bool            `db:"gender"`
	LastName                 string          `db:"last_name"`
	PermanentAddressCommune  string          `db:"permanent_address_commune"`
	PermanentAddressDistrict string          `db:"permanent_address_district"`
	PermanentAddressProvince string          `db:"permanent_address_province"`
	TempAddress              string          `db:"temp_address"`
	AgencyId                 int             `db:"agency_id"`

	FatherBirthPlace  string          `json:"father_birth_place" db:"father_birth_place"`
	MotherBirthPlace  string          `json:"mother_birth_place" db:"mother_birth_place"`
	FatherDob         common.DateTime `json:"father_dob" db:"father_dob" swaggertype:"string"`
	MotherDob         common.DateTime `json:"mother_dob" db:"mother_dob" swaggertype:"string"`
	FatherName        string          `json:"father_name" db:"father_name"`
	MotherName        string          `json:"mother_name" db:"mother_name"`
	Landlord          string          `json:"land_lord" db:"land_lord"`
	FatherOccupation  string          `json:"father_occupation" db:"father_occupation"`
	MotherOccupation  string          `json:"mother_occupation" db:"mother_occupation"`
	FatherPhoneNumber string          `json:"father_phone_number" db:"father_phone_number"`
	MotherPhoneNumber string          `json:"mother_phone_number" db:"mother_phone_number"`
	ResRegistration   string          `json:"parent_res_registration" db:"res_registration"`
	Roi               string          `json:"parent_roi" db:"roi"`
	Zalo              string          `json:"parent_zalo" db:"zalo"`

	ClassCreatedAt time.Time `db:"class_created_at" json:"class_created_at"`
	ClassUpdatedAt time.Time `db:"class_updated_at" json:"class_updated_at"`

	ClassName  string `db:"class_name" json:"class_name"`
	ClassGrade string `db:"class_grade" json:"class_grade"`
	TeacherId  string `db:"class_teacher_id" json:"teacher_id"`
}

type AttendanceResponse struct {
	Attendances []Attendance `json:"attendances"`
	Student     Student      `json:"student"`
}
