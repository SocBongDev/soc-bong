package common

import "time"

type BaseEntity struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindResponse[T any] struct {
	Data []T
	Page int64
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

type Sorter struct {
	Sort   string `query:"sort"   validate:"oneof=asc ASC desc DESC"`
	SortBy string `query:"sortBy"`
}

func (s *Sorter) GetSort() string {
	if s.Sort == "" {
		s.Sort = "DESC"
	}

	return s.Sort
}

func (s *Sorter) GetSortBy() string {
	if s.SortBy == "" {
		s.SortBy = "created_at"
	}

	return s.SortBy
}
