package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type agencyRepo struct {
	db *dbx.DB
}

var _ AgencyRepository = (*agencyRepo)(nil)

func (r *agencyRepo) Find(query *AgencyQuery) ([]Agency, error) {
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

func (r *agencyRepo) FindOne(req *Agency) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *agencyRepo) Insert(req *Agency) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *agencyRepo) Update(req *Agency) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *agencyRepo {
	return &agencyRepo{db}
}
