package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(userPassword, givenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), ([]byte(userPassword)))
	check := true
	if err != nil {
		check = false
		return check
	}
	return check
}
