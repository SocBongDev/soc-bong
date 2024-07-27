package apperr

import "github.com/mdobak/go-xerrors"

var _ error = (*Err)(nil)

type Err struct {
	Err        error
	StatusCode int
}

func (e *Err) Error() string {
	return e.Err.Error()
}

func New(err error) error {
	return xerrors.New(Err{Err: err})
}
