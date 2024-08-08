package classes

import (
	"strings"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/logger"
)

const TABLE = "classes"

var allowedSortFields = map[string]bool{
	"created_at": true,
	"grade":      true,
	"name":       true,
	"updated_at": true,
}

type WriteClassRequest struct {
	Name string `json:"name"`
	// Grade type:
	// * buds - Children who is 3 yo.
	// * seed - Children who is 4 yo.
	// * leaf - Children who is 5 yo.
	Grade     string `json:"grade"     enums:"buds,seed,leaf"`
	AgencyId  int    `json:"agencyId" db:"agency_id"`
	TeacherId string `json:"teacherId" db:"teacher_id"`
}

// ClassQuery represents the query parameters for finding classes
type ClassQuery struct {
	common.Pagination

	AgencyId  int      `json:"agencyId"`
	Ids       []int    `json:"ids"`
	Search    string   `json:"search"`
	TeacherId int      `json:"teacherId"`
	SortBy    []string `json:"sortBy"`
	SortOrder string   `json:"sortOrder"`
}

func (q *ClassQuery) GetOrderBy() string {
	if isSortByEmpty := len(q.SortBy) == 0; isSortByEmpty {
		logger.Info("GetOrderBy", "isSortByEmpty", isSortByEmpty)
		return common.DefaultSortVal
	}

	orderByParts := make([]string, len(q.SortBy))
	for i, field := range q.SortBy {
		if !allowedSortFields[field] {
			logger.Error("GetOrderBy err", "field", field)
			return common.DefaultSortVal
		}

		orderByParts[i] = field
	}
	if q.SortOrder == "" {
		q.SortOrder = common.DefaultSortOrder
	}
	orderByParts[len(q.SortBy)-1] = q.SortOrder

	return strings.Join(orderByParts, ",")
}

type FindClassResp common.FindResponse[Class]

type Class struct {
	common.BaseEntity
	WriteClassRequest
}

func (e *Class) TableName() string {
	return TABLE
}
