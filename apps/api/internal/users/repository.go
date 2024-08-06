package users

type UserRepository interface {
	Find(*UserQuery) ([]User, error)
	FindOne(*User) error
	Insert(*User) error
}
