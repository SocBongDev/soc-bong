package parents

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type WriteParentRequest struct {
	ParentName       string `json:"parent_name"`
	ParentDob        string `json:"parent_dob"`
	ParentGender     bool   `json:"parent_gender"`
	PhoneNumber      string `json:"phone_number"`
	Zalo             string `json:"zalo"`
	Occupation       string `json:"occupation"`
	Landlord         string `json:"landlord"`
	Roi              string `json:"roi"`
	ParentBirthPlace string `json:"parent_birth_place"`
	ResRegistration  string `json:"res_registration"`

	StudentId int `json:"student_id"`
}

type ParentQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type Parent struct {
	common.BaseEntity
	WriteParentRequest
}

var _ = (*dbx.TableModel)(nil)

func (e *Parent) TableName() string {
	return "parents"
}
