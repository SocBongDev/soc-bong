package registrations

import "context"

type RegistrationRepository interface {
	Delete(context.Context, []int) error
	Find(context.Context, *RegistrationQuery) ([]Registration, error)
	FindOne(context.Context, *Registration) error
	Insert(context.Context, *Registration) error
	Update(context.Context, *Registration) error
	MarkAsDone(context.Context, *Registration) error
}
