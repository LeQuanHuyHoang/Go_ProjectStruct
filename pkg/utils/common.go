package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func HashPassword(pass string) (string, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	return string(hashpass), err
}

func CheckPassword(hashpass, pass string) bool {
	rs := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	return rs == nil
}
