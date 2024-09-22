package parents

import "github.com/SocBongDev/soc-bong/internal/common"

const TABLE = "parents"

type WriteParentRequest struct {
	BirthPlace      string `json:"birthPlace"      db:"parent_birth_place"`
	Dob             string `json:"dob"             db:"parent_dob"`
	Gender          bool   `json:"gender"          db:"parent_gender"`
	Landlord        string `json:"landlord"`
	Name            string `json:"name"            db:"parent_name"`
	Occupation      string `json:"occupation"`
	PhoneNumber     string `json:"phoneNumber"`
	ResRegistration string `json:"resRegistration"`
	Roi             string `json:"roi"`
	Zalo            string `json:"zalo"`
	StudentId       int    `json:"-"`
}

type Parent struct {
	common.BaseEntity
	WriteParentRequest
}

func (e *Parent) TableName() string {
	return TABLE
}
