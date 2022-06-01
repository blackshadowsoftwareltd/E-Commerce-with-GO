package helper

import (
	"fmt"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func PasswordToHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("password match ", err == nil)
	return err == nil
}
