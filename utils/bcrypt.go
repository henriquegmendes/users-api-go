package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateEncryptedPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}
