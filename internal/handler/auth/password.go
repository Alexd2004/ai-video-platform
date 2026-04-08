package auth

import (
	"errors"

	"github.com/go-passwd/validator"
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

// https://pkg.go.dev/github.com/go-passwd/validator#ContainsAtLeast used this package
func PasswordValidation(password string) error {

	// The way this api is designed is that commonpassword doesn't validate it itself, but we use it
	// to build a checker, which is why i add the error message i want attached to it, then i set it to a new func
	// Then i can use it
	checkCommon := validator.CommonPassword(errors.New("password is too common"))
	if err := checkCommon(password); err != nil {
		return err
	}

	checkLower := validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, errors.New("must contain a lowercase letter"))
	checkUpper := validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, errors.New("must contain an uppercase letter"))
	checkDigit := validator.ContainsAtLeast("0123456789", 1, errors.New("must contain a number"))
	checkSpecial := validator.ContainsAtLeast("!@#$%^&*()-_=+[]{}|;:,.<>/?~", 1, errors.New("must contain a special character"))

	// Here i am combining all the mini validatir funcs from above into one massive rule set func below
	rules := validator.New(checkLower, checkUpper, checkDigit, checkSpecial)
	// Then i can validate it
	if err := rules.Validate(password); err != nil {
		return err
	}

	return nil
}

// Add tests for these funcs above
// TODO MUST add test
