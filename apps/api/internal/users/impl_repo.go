package users

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type userRepo struct {
	db *dbx.DB
}

func (r *userRepo) Find(query *UserQuery) ([]User, error) {
	resp := make([]User, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("users").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("create_at desc")

	if len(query.Ids) > 0 {
		ids := make([]any, len(query.Ids))
		for i, id := range query.Ids {
			ids[i] = id
		}

		q = q.AndWhere(dbx.In("id", ids...))
	}

	if query.Email != "" {
		q = q.Where(dbx.Like("email", query.Email))
	}

	if query.Search != "" {
		q = q.AndWhere(dbx.Or(dbx.Like("name", query.Search), dbx.Like("grade", query.Search)))
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *userRepo) FindOne(req *User) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *userRepo) Insert(req *User) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func NewRepo(db *dbx.DB) *userRepo {
	return &userRepo{db}
}
