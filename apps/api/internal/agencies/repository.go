package agencies

import "context"

type AgencyRepository interface {
	Find(context.Context, *AgencyQuery) ([]Agency, error)
	FindOne(context.Context, *Agency) error
	Insert(context.Context, *Agency) error
	Update(context.Context, *Agency) error
}
