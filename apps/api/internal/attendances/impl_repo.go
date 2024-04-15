package attendances

import (
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
)

type AttendanceRepo struct {
	db *dbx.DB
}

func (r *AttendanceRepo) Find(query *AttendanceQuery) ([]Attendance, error) {
	resp := []Attendance{}
	q := r.db.Select("*").
		From("attendances").
		Where(dbx.And(
			dbx.HashExp{"class_id": query.ClassId},
			dbx.NewExp(
				"strftime('%Y', attended_at) = {:year}",
				dbx.Params{"year": query.PeriodYear()},
			),
			dbx.NewExp(
				"strftime('%m', attended_at) = {:month}",
				dbx.Params{"month": query.PeriodMonth()},
			),
		))

	if err := q.All(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *AttendanceRepo) Insert(req []Attendance) error {
	vals := make([]string, len(req))
	for i, v := range req {
		vals[i] = fmt.Sprintf(
			"('%s', %d, %t, %d)",
			v.AttendedAt,
			v.ClassId,
			v.IsAttended,
			v.StudentId,
		)
	}

	_, err := r.db.NewQuery(
		fmt.Sprintf(
			`
             INSERT INTO attendances ("attended_at", "class_id", "is_attended", "student_id") 
             VALUES %v
             `,
			strings.Join(vals, ", "),
		),
	).
		Execute()

	return err
}

func (r *AttendanceRepo) Update(req []Attendance) error {
	cases, ids := make([]string, len(req)+1), make([]string, len(req))
	for i, v := range req {
		cases[i] = fmt.Sprintf("WHEN %d THEN %t", v.Id, v.IsAttended)
		ids[i] = fmt.Sprint(v.Id)
	}

	_, err := r.db.NewQuery(
		fmt.Sprintf(
			`
            UPDATE attendances
            SET 'is_attended' = CASE
            %v
            ELSE is_attended
            END
            WHERE id IN (%v)
            `,
			strings.Join(cases, " "),
			strings.Join(ids, ", "),
		),
	).Execute()
	return err
}

func NewRepo(db *dbx.DB) AttendanceRepository {
	return &AttendanceRepo{db}
}
