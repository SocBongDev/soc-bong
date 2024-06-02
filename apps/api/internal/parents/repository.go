package parents

type ParentRepository interface {
	Insert([]*Parent) error
	Update(*Parent) error
}
