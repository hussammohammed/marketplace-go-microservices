package helpers

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
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
	// Concatenate salt with the password
	// log.Println(password)
	// log.Println(string(secretKey))
	// log.Println("****************")
	// passwordWithSalt := append([]byte(password), secretKey...)

	// Hash the password with the generated salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c *CryptHelper) ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
