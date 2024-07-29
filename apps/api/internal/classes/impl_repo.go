package classes

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type classRepo struct {
	db *dbx.DB
}

var _ ClassRepository = (*classRepo)(nil)

func (r *classRepo) Find(query *ClassQuery) ([]*Class, error) {
	resp := make([]*Class, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("classes").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")

	if len(query.Ids) > 0 {
		ids := make([]any, len(query.Ids))
		for i, id := range query.Ids {
			ids[i] = id
		}

		q = q.AndWhere(dbx.In("id", ids...))
	}

	if query.Search != "" {
		q = q.AndWhere(dbx.Or(dbx.Like("name", query.Search), dbx.Like("grade", query.Search)))
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *classRepo) FindOne(req *Class) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *classRepo) Insert(req *Class) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *classRepo) Update(req *Class) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *classRepo {
	return &classRepo{db}
}
