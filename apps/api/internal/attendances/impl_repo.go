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

func (r *AttendanceRepo) Update(req *Attendance) error {
	_, err := r.db.Update(
		"attendances",
		dbx.Params{"is_attended": req.IsAttended},
		dbx.HashExp{"id": req.Id},
	).
		Execute()
	return err
}

func NewRepo(db *dbx.DB) AttendanceRepository {
	return &AttendanceRepo{db}
}
