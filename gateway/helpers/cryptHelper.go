package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

type ICryptHelper interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) error
}

type CryptHelper struct {
}

func NewCryptHelper() *CryptHelper {
	return &CryptHelper{}
}

func (c *CryptHelper) HashPassword(password string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c *CryptHelper) ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
