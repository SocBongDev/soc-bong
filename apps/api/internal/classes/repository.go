package classes

import "context"

type ClassRepository interface {
	Find(context.Context, *ClassQuery) ([]*Class, error)
	FindOne(context.Context, *Class) error
	Insert(context.Context, *Class) error
	Update(context.Context, *Class) error
}
