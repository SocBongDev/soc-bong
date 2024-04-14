package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type RegistrationRepo struct {
	db *dbx.DB
}

func (r *RegistrationRepo) Delete(req *Registration) error {
	return r.db.Model(req).Delete()
}

func (r *RegistrationRepo) Find(query *RegistrationQuery) ([]Registration, error) {
	resp := make([]Registration, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("registrations").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")
	if query.Search != "" {
		q = q.Where(
			dbx.Or(
				dbx.Like("parent_name", query.Search),
				dbx.Like("student_name", query.Search),
				dbx.Like("phone_number", query.Search),
			),
		)
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistrationRepo) FindOne(req *Registration) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *RegistrationRepo) Insert(req *Registration) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *RegistrationRepo) Update(req *Registration) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func (r *RegistrationRepo) MarkAsDone(req *Registration) error {
	_, err := r.db.Update(
		"registrations",
		dbx.Params{"is_processed": true},
		dbx.HashExp{"id": req.Id},
	).
		Execute()
	return err
}

func NewRepo(db *dbx.DB) RegistrationRepository {
	return &RegistrationRepo{db}
}
