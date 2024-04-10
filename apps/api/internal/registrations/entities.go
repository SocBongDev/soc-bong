package registrations

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"gorm.io/datatypes"
)

type Registrations struct {
	common.BaseEntity
	IsProcessed  bool
	Note         string
	ParentName   string
	PhoneNumber  string
	StudentClass string
	StudentDob   datatypes.Date
	StudentName  string
}
