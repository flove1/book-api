package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMismatchedPassword = errors.New("password is incorrent")
)

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CheckPassword(password string, password_hash []byte) error {
	err := bcrypt.CompareHashAndPassword(password_hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return ErrMismatchedPassword
		default:
			return err
		}
	}

	return nil
}
