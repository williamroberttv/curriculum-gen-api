package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 15)
}

func CheckPasswordHash(hash, password string) error {
	fmt.Println(hash)
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}