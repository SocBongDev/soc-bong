package parents

import "github.com/SocBongDev/soc-bong/internal/common"

type WriteParentRequest struct {
	BirthPlace      string `json:"birthPlace"`
	Dob             string `json:"dob"`
	Gender          bool   `json:"gender"`
	Landlord        string `json:"landlord"`
	Name            string `json:"name"`
	Occupation      string `json:"occupation"`
	PhoneNumber     string `json:"phoneNumber"`
	ResRegistration string `json:"resRegistration"`
	Roi             string `json:"roi"`
	Zalo            string `json:"zalo"`
	StudentId       int    `json:"studentId"`
}

type Parent struct {
	common.BaseEntity
	WriteParentRequest
}
