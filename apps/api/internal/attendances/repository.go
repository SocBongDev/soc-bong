package attendances

type AttendanceRepository interface {
	Find(*AttendanceQuery) ([]Attendance, error)
	Insert([]Attendance) error
	Update([]Attendance) error
}
