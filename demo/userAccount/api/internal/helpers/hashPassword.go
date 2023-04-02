package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	// Validate input arguments
	if userPassword == "" || providedPassword == "" {
		return false, "Invalid input arguments"
	}

	// Compare the hashed passwords
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		// Use a constant for the error message
		const errMsg = "Incorrect password"
		return false, errMsg
	}

	// Passwords match
	return true, ""
}
