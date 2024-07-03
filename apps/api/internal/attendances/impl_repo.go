package attendances

import (
	"fmt"
	"strings"

	"github.com/SocBongDev/soc-bong/internal/entities"
	"github.com/pocketbase/dbx"
)

type attendanceRepo struct {
	db *dbx.DB
}

var (
	_            AttendanceRepository = (*attendanceRepo)(nil)
	SelectFields                      = []string{
		fmt.Sprintf("%s.id as id", TABLE),
		fmt.Sprintf("%s.attended_at as attended_at", TABLE),
		fmt.Sprintf("%s.attended_status as attended_status", TABLE),
		fmt.Sprintf("%s.created_at as created_at", TABLE),
		fmt.Sprintf("%s.updated_at as updated_at", TABLE),

		fmt.Sprintf("%s.id as student_id", "students"),
		fmt.Sprintf("%s.created_at as student_created_at", "students"),
		fmt.Sprintf("%s.updated_at as student_updated_at", "students"),
		fmt.Sprintf("%s.birth_place as birth_place", "students"),
		fmt.Sprintf("%s.dob as dob", "students"),
		fmt.Sprintf("%s.enrolled_at as enrolled_at", "students"),
		fmt.Sprintf("%s.ethnic as ethnic", "students"),
		fmt.Sprintf("%s.first_name as first_name", "students"),
		fmt.Sprintf("%s.gender as gender", "students"),
		fmt.Sprintf("%s.last_name as last_name", "students"),
		fmt.Sprintf("%s.permanent_address_commune as permanent_address_commune", "students"),
		fmt.Sprintf("%s.permanent_address_district as permanent_address_district", "students"),
		fmt.Sprintf("%s.permanent_address_province as permanent_address_province", "students"),
		fmt.Sprintf("%s.temp_address as temp_address", "students"),
		fmt.Sprintf("%s.agency_id as agency_id", "students"),
		fmt.Sprintf("%s.id as parent_id", "students"),
		fmt.Sprintf("%s.created_at as parent_created_at", "students"),
		fmt.Sprintf("%s.updated_at as parent_updated_at", "students"),
		fmt.Sprintf("%s.father_birth_place as father_birth_place", "students"),
		fmt.Sprintf("%s.mother_birth_place as mother_birth_place", "students"),
		fmt.Sprintf("%s.father_dob as father_dob", "students"),
		fmt.Sprintf("%s.mother_dob as mother_dob", "students"),
		fmt.Sprintf("%s.mother_name as mother_name", "students"),
		fmt.Sprintf("%s.father_name as father_name", "students"),
		fmt.Sprintf("%s.land_lord as land_lord", "students"),
		fmt.Sprintf("%s.father_occupation as father_occupation", "students"),
		fmt.Sprintf("%s.mother_occupation as mother_occupation", "students"),
		fmt.Sprintf("%s.father_phone_number as father_phone_number", "students"),
		fmt.Sprintf("%s.mother_phone_number as mother_phone_number", "students"),
		fmt.Sprintf("%s.res_registration as res_registration", "students"),
		fmt.Sprintf("%s.roi as roi", "students"),
		fmt.Sprintf("%s.zalo as zalo", "students"),

		fmt.Sprintf("%s.id as class_id", "classes"),
		fmt.Sprintf("%s.created_at as class_created_at", "classes"),
		fmt.Sprintf("%s.updated_at as class_updated_at", "classes"),
		fmt.Sprintf("%s.name as class_name", "classes"),
		fmt.Sprintf("%s.grade as class_grade", "classes"),
		fmt.Sprintf("%s.agency_id as class_agency_id", "classes"),
		fmt.Sprintf("%s.teacher_id as class_teacher_id", "classes"),
	}
)

func (r *attendanceRepo) Find(query *AttendanceQuery) ([]entities.Attendance, error) {
	resp := []entities.Attendance{}
	q := r.db.Select("*").
		From(TABLE).
		Where(
			dbx.And(
				dbx.HashExp{"attendances.class_id": query.ClassId},
				dbx.NewExp(
					"strftime('%Y', attended_at) = {:year}",
					dbx.Params{"year": query.PeriodYear()},
				),
				dbx.NewExp(
					"strftime('%m', attended_at) = {:month}",
					dbx.Params{"month": query.PeriodMonth()},
				),
			))

	if err := q.
		/* InnerJoin("classes", dbx.NewExp(fmt.Sprintf("%s.id = %s.class_id", "classes", TABLE))).
		InnerJoin("students", dbx.NewExp(fmt.Sprintf("%s.id = %s.student_id", "students", TABLE))).
		Build().
		WithAllHook(allHook). */
		All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *attendanceRepo) Insert(req []Attendance) error {
	vals := make([]string, len(req))
	for i, v := range req {
		vals[i] = fmt.Sprintf(
			"('%s', %d, %d, %d)",
			v.AttendedAt,
			v.ClassId,
			v.AttendedStatus,
			v.StudentId,
		)
	}

	_, err := r.db.NewQuery(
		fmt.Sprintf(
			`
             INSERT INTO attendances ("attended_at", "class_id", "attended_status", "student_id") 
             VALUES %v
             `,
			strings.Join(vals, ", "),
		),
	).
		Execute()

	return err
}

func (r *attendanceRepo) Update(req []Attendance) error {
	cases, ids := make([]string, len(req)+1), make([]string, len(req))
	for i, v := range req {
		cases[i] = fmt.Sprintf("WHEN %d THEN %d", v.Id, v.AttendedStatus)
		ids[i] = fmt.Sprint(v.Id)
	}

	_, err := r.db.NewQuery(
		fmt.Sprintf(
			`
            UPDATE attendances
            SET 'attended_status' = CASE
            %v
            ELSE attended_status
            END
            WHERE id IN (%v)
            `,
			strings.Join(cases, " "),
			strings.Join(ids, ", "),
		),
	).Execute()
	return err
}

func NewRepo(db *dbx.DB) *attendanceRepo {
	return &attendanceRepo{db}
}
