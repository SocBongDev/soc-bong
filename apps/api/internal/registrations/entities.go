package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
)

type CreateRegistrationRequest struct {
	Note         string          `json:"note"`
	ParentName   string          `json:"parentName"`
	PhoneNumber  string          `json:"phoneNumber"`
	StudentClass string          `json:"studentClass"`
	StudentDob   common.DateTime `json:"studentDob"`
	StudentName  string          `json:"studentName"`
}

type Registration struct {
	common.BaseEntity
	CreateRegistrationRequest
	IsProcessed bool `json:"isProcessed"`
}

type RegistrationQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}
