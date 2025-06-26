package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {

	passwordHash, err := bcrypt.GenerateFromPassword(password, 11)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func validateHashedPassword(rawPassword string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))

}
