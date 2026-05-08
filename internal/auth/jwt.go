package auth

import (
	"errors"
	"fmt"
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

// ParseAccessToken parses a compact JWT string, verifies HMAC-SHA256 with secret, and checks claims (incl. expiry).
// On success, RegisteredClaims.Subject is the user id issued at login.
func ParseAccessToken(secret, tokenString string) (*jwt.RegisteredClaims, error) {
	if secret == "" {
		return nil, errors.New("jwt secret is empty")
	}
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
