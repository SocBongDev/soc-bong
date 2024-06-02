package classes

import "github.com/SocBongDev/soc-bong/internal/common"

const TABLE = "classes"

type WriteClassRequest struct {
	Name string `json:"name"`
	// Grade type:
	// * buds - Children who is 3 yo.
	// * seed - Children who is 4 yo.
	// * leaf - Children who is 5 yo.
	Grade     string `json:"grade"     enums:"buds,seed,leaf"`
	AgencyId  int    `json:"agencyId"`
	TeacherId string `json:"teacherId"`
}

type ClassQuery struct {
	common.Pagination
	common.Sorter

	AgencyId  int    `json:"agencyId"`
	Search    string `json:"search"`
	TeacherId int    `json:"teacherId"`
}

type FindClassResp common.FindResponse[Class]

type Class struct {
	common.BaseEntity
	WriteClassRequest
}

func (e *Class) TableName() string {
	return TABLE
}
