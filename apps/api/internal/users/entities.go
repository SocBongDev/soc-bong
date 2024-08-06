package users

import (
	"errors"

	"github.com/SocBongDev/soc-bong/internal/common"
)

const TABLE = "users"

type WriteUserRequest struct {
	Email       string          `json:"email"`
	Password    password        `json:"-" db:"password_hash"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	IsActive    bool            `json:"is_active"`
	VerifyEmail bool            `json:"verify_email"`
	Connection  string          `json:"connection"`
	PhoneNumber string          `json:"phone_number"`
	BirthDate   common.DateTime `json:"dob" db:"dob" swaggertype:"string"`
	AgencyId    int             `json:"agencyId" db:"agency_id"`
}

type UserInput struct {
	Email       string          `json:"email" swagger:"required"`
	Password    string          `json:"password" swagger:"required"  db:"password_hash"` // Plaintext password for input
	FirstName   string          `json:"first_name" swagger:"required"`
	LastName    string          `json:"last_name" swagger:"required"`
	IsActive    bool            `json:"is_active"`
	VerifyEmail bool            `json:"verify_email"`
	Connection  string          `json:"connection"`
	PhoneNumber string          `json:"phone_number"`
	BirthDate   common.DateTime `json:"dob" db:"dob" swaggertype:"string"`
	AgencyId    int             `json:"agencyId" db:"agency_id"`
}

type password struct {
	plaintext *string
	hash      string
}

type UserQuery struct {
	common.Pagination
	common.Sorter

	AgencyId int    `json:"agencyId"`
	Ids      []int  `json:"ids"`
	Search   string `json:"search"`
	UserId   string `json:"userId"`
	Email    string `json:"email"`
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type FindUserResp common.FindResponse[User]

type User struct {
	common.BaseEntity
	UserInput
}

func (e *User) TableName() string {
	return TABLE
}
