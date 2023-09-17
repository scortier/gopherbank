package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashedPassword return the bcrypt hash of the password
func HashedPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash password: %w", err)
	}

	return string(pass), nil
}

// CheckPassword checks if the password is correct
func CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("incorrect password: %w", err)
	}

	return nil
}
