package roles

type RoleRepository interface {
	Find(*RoleQuery) ([]Role, error)
	FindOne(*Role) error
	Insert(*Role) error
	Update(*Role) error
}
