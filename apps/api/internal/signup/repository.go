package signup

type UserSignUpRepository interface {
	Insert(*User) error
}
