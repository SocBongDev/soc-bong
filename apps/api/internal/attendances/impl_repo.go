package attendances

import (
	"github.com/SocBongDev/soc-bong/internal/common"
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

func (r *AttendanceRepo) Insert(req *Attendance) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
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
