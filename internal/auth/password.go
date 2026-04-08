package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Need to review the types here
func CheckPasswordhash(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	return err
}

// Make this func nicer using regex, but for now it will be fine. REFACTORs
func PasswordValidation(password string) error {

	if (len(password) < 7) || (len(password) > 20) {
		return errors.New("Password to long/short")
	}

	// Add valdiation

	return nil
}

//Add tests
