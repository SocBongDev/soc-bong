package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
)

type WriteRegistrationRequest struct {
	Note        string `json:"note"`
	ParentName  string `json:"parentName"`
	PhoneNumber string `json:"phoneNumber"`
	// Class type:
	// * buds - Children who is 3 yo.
	// * seed - Children who is 4 yo.
	// * leaf - Children who is 5 yo.
	// * toddler - Children who is 1 - 3 yo.
	StudentClass string          `json:"studentClass" enums:"buds,seed,leaf, toddler"`
	StudentDob   common.DateTime `json:"studentDob"                          swaggertype:"string"`
	StudentName  string          `json:"studentName"`
}

type RegistrationQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type FindRegistrationResp common.FindResponse[Registration]

type DeleteRegistrationQuery struct {
	Ids []int `query:"ids"`
}

type Registration struct {
	common.BaseEntity
	WriteRegistrationRequest
	IsProcessed bool `json:"isProcessed"`
}

func (e *Registration) TableName() string {
	return "registrations"
}
