package classes

import "github.com/SocBongDev/soc-bong/internal/common"

const TABLE = "classes"

type WriteClassRequest struct {
	Name string `json:"name"`
	// Grade type:
	// * buds - Children who is 3 yo.
	// * seed - Children who is 4 yo.
	// * leaf - Children who is 5 yo.
	// * toddler - Children who is lower than 2 yo.
	Grade     string `json:"grade"     enums:"buds,seed,leaf, toddler"`
	AgencyId  int    `json:"agencyId" db:"agency_id"`
	TeacherId string `json:"teacherId" db:"teacher_id"`
}

type ClassQuery struct {
	common.Pagination
	common.Sorter

	AgencyId  int    `json:"agencyId"`
	Ids       []int  `json:"ids"`
	Search    string `json:"search"`
	TeacherId string `json:"teacherId"`
}

type FindClassResp common.FindResponse[*Class]

type Class struct {
	common.BaseEntity
	WriteClassRequest
}

func (e *Class) TableName() string {
	return TABLE
}
