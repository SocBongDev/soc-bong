package students

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type studentRepo struct {
	db *dbx.DB
}

var (
	_            StudentRepository = (*studentRepo)(nil)
	SelectFields                   = []string{
		fmt.Sprintf("%s.id as id", TABLE),
		fmt.Sprintf("%s.created_at as created_at", TABLE),
		fmt.Sprintf("%s.updated_at as updated_at", TABLE),
		fmt.Sprintf("%s.birth_place as birth_place", TABLE),
		fmt.Sprintf("%s.dob as dob", TABLE),
		fmt.Sprintf("%s.enrolled_at as enrolled_at", TABLE),
		fmt.Sprintf("%s.ethnic as ethnic", TABLE),
		fmt.Sprintf("%s.first_name as first_name", TABLE),
		fmt.Sprintf("%s.gender as gender", TABLE),
		fmt.Sprintf("%s.last_name as last_name", TABLE),
		fmt.Sprintf("%s.permanent_address_commune as permanent_address_commune", TABLE),
		fmt.Sprintf("%s.permanent_address_district as permanent_address_district", TABLE),
		fmt.Sprintf("%s.permanent_address_province as permanent_address_province", TABLE),
		fmt.Sprintf("%s.temp_address as temp_address", TABLE),
		fmt.Sprintf("%s.agency_id as agency_id", TABLE),
		fmt.Sprintf("%s.id as class_id", classes.TABLE),
		fmt.Sprintf("%s.created_at as class_created_at", classes.TABLE),
		fmt.Sprintf("%s.updated_at as class_updated_at", classes.TABLE),
		fmt.Sprintf("%s.name as class_name", classes.TABLE),
		fmt.Sprintf("%s.grade as class_grade", classes.TABLE),
		fmt.Sprintf("%s.agency_id as class_agency_id", classes.TABLE),
		fmt.Sprintf("%s.teacher_id as class_teacher_id", classes.TABLE),
		fmt.Sprintf("%s.id as parent_id", TABLE),
		fmt.Sprintf("%s.created_at as parent_created_at", TABLE),
		fmt.Sprintf("%s.updated_at as parent_updated_at", TABLE),
		fmt.Sprintf("%s.father_birth_place as father_birth_place", TABLE),
		fmt.Sprintf("%s.mother_birth_place as mother_birth_place", TABLE),
		fmt.Sprintf("%s.father_dob as father_dob", TABLE),
		fmt.Sprintf("%s.mother_dob as mother_dob", TABLE),
		fmt.Sprintf("%s.mother_name as mother_name", TABLE),
		fmt.Sprintf("%s.father_name as father_name", TABLE),
		fmt.Sprintf("%s.land_lord as land_lord", TABLE),
		fmt.Sprintf("%s.father_occupation as father_occupation", TABLE),
		fmt.Sprintf("%s.mother_occupation as mother_occupation", TABLE),
		fmt.Sprintf("%s.father_phone_number as father_phone_number", TABLE),
		fmt.Sprintf("%s.mother_phone_number as mother_phone_number", TABLE),
		fmt.Sprintf("%s.res_registration as res_registration", TABLE),
		fmt.Sprintf("%s.roi as roi", TABLE),
		fmt.Sprintf("%s.zalo as zalo", TABLE),
	}
)

func (r *studentRepo) Delete(ids []int) error {
	anySlices := make([]any, len(ids))
	for i, v := range ids {
		anySlices[i] = v
	}

	_, err := r.db.Delete("students", dbx.HashExp{"id": anySlices}).Execute()
	return err
}

func (r *studentRepo) Find(query *StudentQuery) ([]Student, error) {
	resp := make([]Student, 0, query.GetPageSize())
	q := r.db.Select(SelectFields...).
		From(TABLE).
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

	if err := q.
		InnerJoin(
			classes.TABLE,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.id = %s.class_id",
					classes.TABLE,
					TABLE,
				),
			),
		).
		Build().
		WithAllHook(allHook).
		All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *studentRepo) FindOne(req *Student) error {
	students := []Student{}
	if err := r.db.Select(SelectFields...).From(TABLE).
		Where(dbx.HashExp{fmt.Sprintf("%s.id", TABLE): req.Id}).
		InnerJoin(
			classes.TABLE,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.id = %s.class_id",
					classes.TABLE,
					TABLE,
				),
			),
		).
		Build().
		WithAllHook(allHook).
		All(&students); err != nil {
		return err
	}
	if len(students) == 0 {
		return sql.ErrNoRows
	}

	if len(students) == 1 {
		*req = students[0]
	}

	return nil
}

func (r *studentRepo) Insert(req *Student) error {
	log.Printf("check final req: %v", req)
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *studentRepo) Update(req *Student) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *studentRepo {
	return &studentRepo{db}
}
