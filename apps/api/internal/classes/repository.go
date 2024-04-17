package classes

type ClassRepository interface {
	Find(*ClassQuery) ([]Class, error)
	FindOne(*Class) error
	Insert(*Class) error
	Update(*Class) error
}
