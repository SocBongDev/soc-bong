package students

import "github.com/SocBongDev/soc-bong/internal/entities"

type StudentRepository interface {
	Delete([]int) error
	Find(*StudentQuery) ([]entities.Student, error)
	FindOne(*Student) error
	Insert(*Student) error
	Update(*Student) error
}
