package users

import "context"

type Repository interface {
	Find(context.Context, *UserQuery) ([]User, error)
	FindOne(context.Context, *User) error
	Insert(context.Context, *User) error
}
