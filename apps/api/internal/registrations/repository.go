package registrations

type RegistrationRepository interface {
	Delete([]int) error
	Find(*RegistrationQuery) ([]Registration, error)
	FindOne(*Registration) error
	Insert(*Registration) error
	Update(*Registration) error
	MarkAsDone(*Registration) error
}
