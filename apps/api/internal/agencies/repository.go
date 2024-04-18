package agencies

type AgencyRepository interface {
	Find(*AgencyQuery) ([]Agency, error)
	FindOne(*Agency) error
	Insert(*Agency) error
	Update(*Agency) error
}
