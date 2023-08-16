package format

import "net/mail"

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsPasswordValid(password string) bool {
	lengthIsOk := len(password) >= 8
	
	return lengthIsOk
}