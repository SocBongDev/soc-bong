package registrations

type RegistrationRepository interface {
	Delete(*Registration) error
	Find(*RegistrationQuery) ([]Registration, error)
	FindOne(*Registration) error
	Insert(*Registration) error
	Update(*Registration) error
}
