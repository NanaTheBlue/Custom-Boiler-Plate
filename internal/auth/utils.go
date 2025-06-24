package auth

import (
	"errors"
	"regexp"

	"github.com/nanagoboiler/models"
	"golang.org/x/crypto/bcrypt"
)

func validateUsername(username string) error {
	var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,15}$`)

	if len(username) > 15 || len(username) < 3 {
		return errors.New("Invalid Username Length")

	}

	if !usernameRegex.MatchString(username) {
		return errors.New("Invalid Characters in Username")
	}

	return nil

}

func validatePassword(password string, confirmpassword string) error {

	// Just basic password validation for now
	if len(password) < 8 {
		return errors.New("Password Length Is To Short")
	}
	if password != confirmpassword {
		return errors.New("Passwords Dont Match")
	}

	return nil
}

func validateEmail(email string) error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid Email Format")
	}

	return nil

}

func validateRegistration(req *models.RegisterRequest) error {

	err := validatePassword(req.Password, req.ConfirmPassword)
	if err != nil {
		return err
	}
	err = validateEmail(req.Email)
	if err != nil {
		return err
	}
	err = validateUsername(req.Username)
	if err != nil {
		return err
	}

	return nil
}

func ValidateLogin(req *models.LoginRequest) error {

	err := validateEmail(req.Email)
	if err != nil {
		return err
	}

	return nil
}

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
