package attendances

import (
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
)

type attendanceRepo struct {
	db *dbx.DB
}

var _ AttendanceRepository = (*attendanceRepo)(nil)

func (r *attendanceRepo) Find(query *AttendanceQuery) ([]Attendance, error) {
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
