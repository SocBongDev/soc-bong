package parents

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type ParentRepo struct {
	db *dbx.DB
}

func (r *ParentRepo) Delete(req *Parent) error {
	return r.db.Model(req).Delete()
}

func (r *ParentRepo) Find(query *ParentQuery) ([]Parent, error) {
	resp := make([]Parent, 0, query.GetPageSize())
	q := r.db.Select("*").
		From("parents").
		Offset(query.GetOffset()).
		Limit(query.GetPageSize()).
		OrderBy("created_at desc")
	if query.Search != "" {
		q = q.Where(
			dbx.Or(
				dbx.Like("parent_name", query.Search),
				dbx.Like("phone_number", query.Search),
			),
		)
	}

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ParentRepo) FindOne(req *Parent) error {
	return r.db.Select().Model(req.Id, req)
}

func (r *ParentRepo) Insert(req *Parent) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *ParentRepo) Update(req *Parent) error {
	if req.StudentId != 0 {

		params := dbx.Params{
			"parent_name":        req.ParentName,
			"parent_dob":         req.ParentDob,
			"parent_gender":      req.ParentGender,
			"phone_number":       req.PhoneNumber,
			"zalo":               req.Zalo,
			"occupation":         req.Occupation,
			"landlord":           req.Landlord,
			"roi":                req.Roi,
			"parent_birth_place": req.ParentBirthPlace,
			"res_registration":   req.ResRegistration,
		}

		_, err := r.db.Update(
			"parents",
			params,
			dbx.HashExp{"studentId": req.StudentId}).Execute()

		return err
	}

	return error(nil)
}

func NewRepo(db *dbx.DB) ParentRepository {
	return &ParentRepo{db}
}
