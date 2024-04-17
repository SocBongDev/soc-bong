package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type WriteStudentRequest struct {
	FirstName                string          `json:"first_name"`
	LastName                 string          `json:"last_name"`
	EnrolledAt               common.DateTime `json:"enrolled_at" swaggertype:"string"`
	Dob                      common.DateTime `json:"dob" swaggertype:"string"`
	Gender                   bool            `json:"gender"`
	Ethnic                   string          `json:"ethnic"`
	BirthPlace               string          `json:"birth_place"`
	TempAddress              string          `json:"temp_address"`
	PermanentAddressProvince string          `json:"permanent_address_province"`
	PermanentAddressDistrict string          `json:"permanent_address_district"`
	PermanentAddressCommune  string          `json:"permanent_address_commune"`
	ClassId                  int             `json:"class_id"`
	AgencyId                 int             `json:"agency_id"`
}

type StudentQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type Student struct {
	common.BaseEntity
	WriteStudentRequest
}

var _ = (*dbx.TableModel)(nil)

func (e *Student) TableName() string {
	return "students"
}
