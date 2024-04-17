package parents

type ParentRepository interface {
	Delete(*Parent) error
	Find(*ParentQuery) ([]Parent, error)
	FindOne(*Parent) error
	Insert(*Parent) error
	Update(*Parent) error
}
