package classes

import "github.com/SocBongDev/soc-bong/internal/common"

type WriteClassRequest struct {
	Name string `json:"name"`
	// Grade type:
	// * buds - Children who is 3 yo.
	// * seed - Children who is 4 yo.
	// * leaf - Children who is 5 yo.
	Grade     string `json:"grade"      enums:"buds,seed,leaf"`
	AgencyId  int    `json:"agency_id"`
	TeacherId string `json:"teacher_id"`
}

type ClassQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type FindClassResp common.FindResponse[Class]

type Class struct {
	common.BaseEntity
	WriteClassRequest
}

func (e *Class) TableName() string {
	return "classes"
}
