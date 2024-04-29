package students

import "github.com/SocBongDev/soc-bong/internal/common"

type WriteStudentRequest struct {
	BirthPlace               string          `json:"birthPlace"`
	Dob                      common.DateTime `json:"dob"`
	EnrolledAt               common.DateTime `json:"enrolledAt"`
	Ethnic                   string          `json:"ethnic"`
	FirstName                string          `json:"firstName"`
	Gender                   bool            `json:"gender"`
	LastName                 string          `json:"lastName"`
	PermanentAddressCommune  string          `json:"permanentAddressCommune"`
	PermanentAddressDistrict string          `json:"permanentAddressDistrict"`
	PermanentAddressProvince string          `json:"permanentAddressProvince"`
	TempAdress               string          `json:"tempAdress"`
	AgencyId                 int             `json:"agencyId"`
	ClassId                  int             `json:"classId"`
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

type Student struct {
	common.BaseEntity
	WriteStudentRequest
}

func (e *Student) TableName() string {
	return "students"
}
