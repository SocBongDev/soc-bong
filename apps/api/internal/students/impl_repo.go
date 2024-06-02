package students

import (
	"database/sql"
	"fmt"

	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/parents"
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
		fmt.Sprintf("%s.id as parent_id", parents.TABLE),
		fmt.Sprintf("%s.created_at as parent_created_at", parents.TABLE),
		fmt.Sprintf("%s.updated_at as parent_updated_at", parents.TABLE),
		fmt.Sprintf("%s.parent_birth_place as parent_birth_place", parents.TABLE),
		fmt.Sprintf("%s.parent_dob as parent_dob", parents.TABLE),
		fmt.Sprintf("%s.parent_gender as parent_gender", parents.TABLE),
		fmt.Sprintf("%s.landlord as parent_landlord", parents.TABLE),
		fmt.Sprintf("%s.parent_name as parent_name", parents.TABLE),
		fmt.Sprintf("%s.occupation as parent_occupation", parents.TABLE),
		fmt.Sprintf("%s.phone_number as parent_phone_number", parents.TABLE),
		fmt.Sprintf("%s.res_registration as parent_res_registration", parents.TABLE),
		fmt.Sprintf("%s.roi as parent_roi", parents.TABLE),
		fmt.Sprintf("%s.zalo as parent_zalo", parents.TABLE),
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
		InnerJoin(
			parents.TABLE,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.student_id = %s.id",
					parents.TABLE,
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
		InnerJoin(
			parents.TABLE,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.student_id = %s.id",
					parents.TABLE,
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
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *studentRepo) Update(req *Student) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *studentRepo {
	return &studentRepo{db}
}
