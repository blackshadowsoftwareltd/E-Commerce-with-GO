package helper

import "net/mail"

func IsEmailValid(email string) bool {
	_, _emailErr := mail.ParseAddress(email)
	if _emailErr != nil {

		return false
	} else {
		return true
	}
}
func IsPasswordValid(password string) bool {
	if password == "" || len(password) < 6 {
		return false
	} else {
		return true
	}
}
