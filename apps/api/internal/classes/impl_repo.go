package classes

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type ClassRepo struct {
	db *dbx.DB
}

func (r *ClassRepo) Find(query *ClassQuery) ([]Class, error) {
	resp := make([]Class, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("classes").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")
	if query.Search != "" {
		q = q.Where(dbx.Or(dbx.Like("name", query.Search), dbx.Like("grade", query.Search)))
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ClassRepo) FindOne(req *Class) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *ClassRepo) Insert(req *Class) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *ClassRepo) Update(req *Class) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) ClassRepository {
	return &ClassRepo{db}
}
