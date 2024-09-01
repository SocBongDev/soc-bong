package classes

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type classRepo struct {
	db *dbx.DB
}

var _ ClassRepository = (*classRepo)(nil)

func (r *classRepo) Find(ctx context.Context, query *ClassQuery) ([]*Class, error) {
	resp := make([]*Class, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("classes").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy(query.GetOrderBy())

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

	if query.TeacherId != "" {
		q = q.AndWhere(dbx.NewExp("teacher_id = {:teacher_id}", dbx.Params{"teacher_id": query.TeacherId}))
	}

	if err := q.WithContext(ctx).All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *classRepo) FindOne(ctx context.Context, req *Class) error {
	return r.db.WithContext(ctx).Select().Model(req.Id, req)
}

func (r *classRepo) Insert(ctx context.Context, req *Class) error {
	return r.db.WithContext(ctx).Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *classRepo) Update(ctx context.Context, req *Class) error {
	return r.db.WithContext(ctx).Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *classRepo {
	return &classRepo{db}
}
