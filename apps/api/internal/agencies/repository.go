package agencies

type AgencyRepository interface {
	Delete([]int) error
	Find(*AgencyQuery) ([]Agency, error)
	FindOne(*Agency) error
	Insert(*Agency) error
	Update(*Agency) error
	MarkAsDone(*Agency) error
}
