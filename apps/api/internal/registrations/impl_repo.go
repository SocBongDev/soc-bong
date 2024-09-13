package registrations

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type registrationRepo struct {
	db *dbx.DB
}

var _ RegistrationRepository = (*registrationRepo)(nil)

func (r *registrationRepo) Delete(ctx context.Context, req []int) error {
	anySlices := make([]any, len(req))
	for i, v := range req {
		anySlices[i] = v
	}

	_, err := r.db.WithContext(ctx).Delete("registrations", dbx.HashExp{"id": anySlices}).Execute()
	return err
}

func (r *registrationRepo) Find(ctx context.Context, query *RegistrationQuery) ([]Registration, error) {
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
				dbx.Like("agency_id", query.Search),
			),
		)
	}

	if err := q.WithContext(ctx).All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *registrationRepo) FindOne(ctx context.Context, req *Registration) error {
	return r.db.WithContext(ctx).Select().Model(req.Id, req)
}

func (r *registrationRepo) Insert(ctx context.Context, req *Registration) error {
	return r.db.WithContext(ctx).Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *registrationRepo) Update(ctx context.Context, req *Registration) error {
	return r.db.WithContext(ctx).Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func (r *registrationRepo) MarkAsDone(ctx context.Context, req *Registration) error {
	_, err := r.db.WithContext(ctx).
		Update(
			"registrations",
			dbx.Params{"is_processed": true},
			dbx.HashExp{"id": req.Id},
		).
		Execute()
	return err
}

func NewRepo(db *dbx.DB) *registrationRepo {
	return &registrationRepo{db}
}
