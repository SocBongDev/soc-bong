package students

import (
	"context"

	"github.com/SocBongDev/soc-bong/internal/entities"
)

type StudentRepository interface {
	Delete(context.Context, []int) error
	Find(context.Context, *StudentQuery) ([]entities.Student, error)
	FindOne(context.Context, *Student) error
	Insert(context.Context, *Student) error
	Update(context.Context, *Student) error
}
