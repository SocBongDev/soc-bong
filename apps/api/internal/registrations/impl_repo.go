package registrations

import "github.com/pocketbase/dbx"

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
		OrderBy("created desc")
	if query.Search != "" {
		q = q.Where(
			dbx.Or(
				dbx.Like("parentName"),
				dbx.Like("studentName"),
				dbx.Like("phoneNumber"),
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
	return r.db.Model(req).Insert()
}

func (r *RegistrationRepo) Update(req *Registration) error {
	return r.db.Model(req).Update()
}

func NewRepo(db *dbx.DB) RegistrationRepository {
	return &RegistrationRepo{db}
}
