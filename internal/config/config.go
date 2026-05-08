package config

import (
	"errors"
	"os"
)

// Config holds process-wide settings read from the environment.
// Set these in the shell or your process manager, look at the env e xam[le]
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	// JWTSecret signs access tokens (login). Must be long and random in production. // REVIVEW
	JWTSecret string
	// HTTPAddr is the listen address for Gin, e.g. ":8080" or "127.0.0.1:8080". // REVIEW
	HTTPAddr string
}

// Load reads configuration from environment variables.
// Returns an error if JWT_SECRET is missing, we need this as the apis fail without it
func Load() (Config, error) {
	c := Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		HTTPAddr:   os.Getenv("HTTP_ADDR"),
	}
	if c.HTTPAddr == "" {
		c.HTTPAddr = ":8080"
	}
	if c.JWTSecret == "" {
		return c, errors.New("JWT_SECRET is required")
	}
	return c, nil
}
