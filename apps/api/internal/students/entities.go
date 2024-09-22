package students

import (
	"time"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/entities"
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
	AgencyId                 int             `json:"agencyId" db:"agency_id"`
	ClassId                  int             `json:"classId" db:"class_id"`

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
}

type WriteParentRequest struct {
	FatherBirthPlace  string          `json:"father_birth_place" db:"father_birth_place"`
	MotherBirthPlace  string          `json:"mother_birth_place" db:"mother_birth_place"`
	FatherDob         common.DateTime `json:"father_dob" db:"father_dob" swaggertype:"string"`
	MotherDob         common.DateTime `json:"mother_dob" db:"mother_dob" swaggertype:"string"`
	FatherName        string          `json:"father_name" db:"father_name"`
	MotherName        string          `json:"mother_name" db:"mother_name"`
	Landlord          string          `json:"parent_land_lord" db:"land_lord"`
	FatherOccupation  string          `json:"father_occupation" db:"father_occupation"`
	MotherOccupation  string          `json:"mother_occupation" db:"mother_occupation"`
	FatherPhoneNumber string          `json:"father_phone_number" db:"father_phone_number"`
	MotherPhoneNumber string          `json:"mother_phone_number" db:"mother_phone_number"`
	ResRegistration   string          `json:"parent_res_registration" db:"res_registration"`
	Roi               string          `json:"parent_roi" db:"roi"`
	Zalo              string          `json:"parent_zalo" db:"zalo"`
}

type InsertStudentRequest struct {
	WriteStudentRequest
	Parents []WriteParentRequest `json:"parents"`
}

type Parent struct {
	common.BaseEntity
	WriteParentRequest
}

type StudentQuery struct {
	common.Pagination
	common.Sorter

	ClassId int    `json:"classId"`
	Ids     []int  `json:"ids"`
	Search  string `json:"search"`
}

type FindStudentResp common.FindResponse[entities.Student]

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
	AgencyId   int    `db:"class_agency_id"`
	TeacherId  string `db:"class_teacher_id"`
}

type Student struct {
	common.BaseEntity
	WriteStudentRequest

	Class entities.Class `json:"class"   db:"-"`
}

func (e *Student) TableName() string {
	return TABLE
}
