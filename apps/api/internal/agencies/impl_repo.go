package agencies

import (
	"fmt"
	"log"

	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type AgencyRepo struct {
	db *dbx.DB
}

func (r *AgencyRepo) Delete(req []int) error {
	idsSlices := make([]any, len(req))
	for i, v := range req {
		idsSlices[i] = v
	}

	var studentsCount int

	err := r.db.Select("COUNT(*)").From("students").Where(dbx.HashExp{"agency_id": idsSlices}).Row(&studentsCount)
	if err != nil {
		log.Println("AgencyRepo.Delete.Select students err: ", err)
		return err
	}

	var parentsCount int

	err = r.db.Select("COUNT(*)").From("parents").Where(dbx.HashExp{"agency_id": idsSlices}).Row(&parentsCount)
	if err != nil {
		log.Println("AgencyRepo.Delete.Select parents err: ", err)
		return err
	}

	if studentsCount > 0 || parentsCount > 0 {
		return fmt.Errorf("AgencyRepo.Delete.err: agency is still in use")
	}

	_, err = r.db.Delete("agencies", dbx.HashExp{"id": idsSlices}).Execute()

	if err != nil {
		log.Println("AgencyRepo.Delete.err: ", err)
		return err
	}
	return nil
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
				dbx.Like("name", query.Search),
				dbx.Like("phone_number", query.Search),
				dbx.Like("email", query.Search),
				dbx.Like("address", query.Search),
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

func (r *AgencyRepo) MarkAsDone(req *Agency) error {
	_, err := r.db.Update(
		"agencies",
		dbx.Params{"is_processed": true},
		dbx.HashExp{"id": req.Id},
	).
		Execute()
	return err
}

func NewRepo(db *dbx.DB) AgencyRepository {
	return &AgencyRepo{db}
}
