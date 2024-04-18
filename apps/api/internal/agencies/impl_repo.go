package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type AgencyRepo struct {
	db *dbx.DB
}

func (r *AgencyRepo) Find(query *AgencyQuery) ([]Agency, error) {
	resp := make([]Agency, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("agencies").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")
	if query.Search != "" {
		q = q.Where(
			dbx.Or(
				dbx.Like("address", query.Search),
				dbx.Like("email", query.Search),
				dbx.Like("name", query.Search),
				dbx.Like("phone", query.Search),
			),
		)
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *AgencyRepo) FindOne(req *Agency) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *AgencyRepo) Insert(req *Agency) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *AgencyRepo) Update(req *Agency) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) AgencyRepository {
	return &AgencyRepo{db}
}
