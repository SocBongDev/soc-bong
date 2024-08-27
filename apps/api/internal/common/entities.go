package common

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type BaseEntity struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FindResponse[T any] struct {
	Data     []T   `json:"data"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type Pagination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

func (p *Pagination) GetPage() int64 {
	if p.Page == 0 || p.Page < 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetPageSize() int64 {
	switch {
	case p.PageSize <= 0:
		p.PageSize = 10
	case p.PageSize >= 100:
		p.PageSize = 100
	}
	return p.PageSize
}

func (p *Pagination) GetOffset() int64 {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// SortField represents a field to sort by and its order
type SortField struct {
	// @Description The field name to sort by
	// @Example name
	Field string `json:"field"`

	// @Description The sort order (asc or desc)
	// @Enum asc,ASC,desc,DESC
	// @Example ASC
	Order string `json:"order" validate:"oneof=asc ASC desc DESC"`
}

type Sorter struct {
	// @Description List of fields to sort by
	SortFields []SortField `json:"sortFields"`
}

var BaseExcludeFields []string = []string{"Id", "CreatedAt", "UpdatedAt"}

var (
	DefaultSortVal   = "created_at DESC"
	DefaultSortOrder = "DESC"
)

type APIHandler interface {
	RegisterRoute(fiber.Router)
}
