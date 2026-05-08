package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: REVIEW THIS FILE HARD
// defaultTokenTTL is how long access tokens remain valid unless you add refresh tokens later.
const defaultTokenTTL = 24 * time.Hour

// SignAccessToken builds a signed JWT using HMAC-SHA256.
// Claims use RegisteredClaims: Subject is the user id (UUID string from the database).
// The client can send the token in Authorization: Bearer <token> on protected routes.
func SignAccessToken(secret, userID string) (string, error) {
	if secret == "" {
		return "", errors.New("jwt secret is empty")
	}
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(defaultTokenTTL)),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}
