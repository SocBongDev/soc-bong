package students

import (
	"time"

	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/parents"
)

const TABLE = "students"

type WriteStudentRequest struct {
	BirthPlace               string          `json:"birthPlace"`
	Dob                      common.DateTime `json:"dob"                      swaggertype:"string"`
	EnrolledAt               common.DateTime `json:"enrolledAt"               swaggertype:"string"`
	Ethnic                   string          `json:"ethnic"`
	FirstName                string          `json:"firstName"`
	Gender                   bool            `json:"gender"`
	LastName                 string          `json:"lastName"`
	PermanentAddressCommune  string          `json:"permanentAddressCommune"`
	PermanentAddressDistrict string          `json:"permanentAddressDistrict"`
	PermanentAddressProvince string          `json:"permanentAddressProvince"`
	TempAddress              string          `json:"tempAddress"`
	AgencyId                 int             `json:"agencyId"`
	ClassId                  int             `json:"classId"`
}

type InsertStudentRequest struct {
	WriteStudentRequest
	Parents []parents.WriteParentRequest `json:"parents"`
}

type StudentQuery struct {
	common.Pagination
	common.Sorter
	ClassId int    `json:"classId"`
	Search  string `json:"search"`
}

type FindStudentResp common.FindResponse[Student]

type DeleteStudentQuery struct {
	Ids []int `query:"ids"`
}

type DbStudentResult struct {
	common.BaseEntity
	WriteStudentRequest

	ClassId        int       `db:"class_id"`
	ClassCreatedAt time.Time `db:"class_created_at"`
	ClassUpdatedAt time.Time `db:"class_updated_at"`

	ClassName  string `db:"class_name"`
	ClassGrade string `db:"class_grade"`

	ParentId        int       `db:"parent_id"`
	ParentCreatedAt time.Time `db:"parent_created_at"`
	ParentUpdatedAt time.Time `db:"parent_updated_at"`

	ParentBirthPlace      string `db:"parent_birth_place"`
	ParentDob             string `db:"parent_dob"`
	ParentGender          bool   `db:"parent_gender"`
	ParentLandlord        string `db:"parent_landlord"`
	ParentName            string `db:"parent_name"`
	ParentOccupation      string `db:"parent_occupation"`
	ParentPhoneNumber     string `db:"parent_phone_number"`
	ParentResRegistration string `db:"parent_res_registration"`
	ParentRoi             string `db:"parent_roi"`
	ParentZalo            string `db:"parent_zalo"`
}

type Student struct {
	common.BaseEntity
	WriteStudentRequest

	Class   classes.Class    `json:"class"   db:"-"`
	Parents []parents.Parent `json:"parents" db:"-"`
}

func (e *Student) TableName() string {
	return TABLE
}
