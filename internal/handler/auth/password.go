package auth

import (
	"errors"

	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
)

// bcryptCost is the work factor for bcrypt.GenerateFromPassword (higher = slower and harder to brute-force).
const bcryptCost = 14

// PasswordHash returns a bcrypt hash suitable for storing in users.password_hash.
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password to a stored bcrypt hash
// Returns nil if they match; otherwise an error from bcrypt
func CheckPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// PasswordValidation runs policy checks before hashing (registration).
// Uses github.com/go-passwd/validator: common password list plus characters
//
// Se  https://pkg.go.dev/github.com/go-passwd/validator#ContainsAtLeast
func PasswordValidation(password string) error {
	// CommonPassword returns a ValidateFunc; pass a custom error to return when the password is in the list
	checkCommon := validator.CommonPassword(errors.New("password is too common"))
	if err := checkCommon(password); err != nil {
		return err
	}

	checkLower := validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, errors.New("must contain a lowercase letter"))
	checkUpper := validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, errors.New("must contain an uppercase letter"))
	checkDigit := validator.ContainsAtLeast("0123456789", 1, errors.New("must contain a number"))
	checkSpecial := validator.ContainsAtLeast("!@#$%^&*()-_=+[]{}|;:,.<>/?~", 1, errors.New("must contain a special character"))

	rules := validator.New(checkLower, checkUpper, checkDigit, checkSpecial)
	if err := rules.Validate(password); err != nil {
		return err
	}

	return nil
}
