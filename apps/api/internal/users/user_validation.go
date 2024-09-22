package users

import "github.com/SocBongDev/soc-bong/internal/common"

func ValidateEmail(v *common.Validator, email string) {
	v.Check(email != "", "email", "email must be provided")
	v.Check(common.Matches(email, common.EmailRX), "email", "email must be a valid email address")
}

func ValidatePasswordPlaintext(v *common.Validator, password string) {
	v.Check(password != "", "password", "password must be provided")
	v.Check(len(password) >= 8, "password", "password must be at least 8 bytes")
	v.Check(len(password) <= 72, "password", "password not be more than 72 bytes long")
}

func ValidateUser(v *common.Validator, user *User) {
	v.Check(user.FirstName != "", "first_name", "first name must be provided")
	v.Check(user.LastName != "", "last_name", "last name must be provided")

	ValidateEmail(v, user.Email)
	//If the plaintext password is not nil, call the standalone //ValidatePasswordPlaintext() helper.

	if user.Password != "" {
		ValidatePasswordPlaintext(v, user.Password)
	}
}
