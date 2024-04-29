package students

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type StudentRepo struct {
	db *dbx.DB
}

func (r *StudentRepo) Delete(ids []int) error {
	anySlices := make([]any, len(ids))
	for i, v := range ids {
		anySlices[i] = v
	}

	_, err := r.db.Delete("students", dbx.HashExp{"id": anySlices}).Execute()
	return err
}

func (r *StudentRepo) Find(query *StudentQuery) ([]Student, error) {
	resp := make([]Student, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("students").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")
	if query.Search != "" {
		q = q.Where(
			dbx.Or(
				dbx.Like("first_name", query.Search),
				dbx.Like("last_name", query.Search),
			),
		)
	}
	if query.ClassId != 0 {
		q = q.AndWhere(dbx.HashExp{"class_id": query.ClassId})
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *StudentRepo) FindOne(req *Student) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *StudentRepo) Insert(req *Student) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *StudentRepo) Update(req *Student) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) StudentRepository {
	return &StudentRepo{db}
}
