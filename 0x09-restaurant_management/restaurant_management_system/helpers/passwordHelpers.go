package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
        log.Panic(err)
    }
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	var msg string = "credentials verified successfully"
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			msg = "incorrect credentials"
            return false, msg
        }
        log.Fatal(err)
		msg = "error verifying credentials"
        return false, msg
	}
	return true, msg
}