package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (p *password) Set(plaintextPassword string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if error != nil {
		log.Println("error hashing: ", error)
		return "", error
	}

	p.plaintext = &plaintextPassword
	p.hash = string(hash)
	return p.hash, nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
