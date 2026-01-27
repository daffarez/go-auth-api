package utils

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" || !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func ValidateRegisterInput(email, password string) error {
	if err := ValidateEmail(email); err != nil {
		return err
	}
	if err := ValidatePassword(password); err != nil {
		return err
	}
	return nil
}
