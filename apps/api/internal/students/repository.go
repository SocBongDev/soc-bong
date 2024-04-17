package students

type StudentRepository interface {
	Delete(*Student) error
	Find(*StudentQuery) ([]Student, error)
	FindOne(*Student) error
	Insert(*Student) error
	Update(*Student) error
}
