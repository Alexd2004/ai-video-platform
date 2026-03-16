package database

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	dbUser, dbPassword, dbHost, dbPort, dbName := os.Getenv(("DB_USER")), os.Getenv(("DB_PASSWORD")), os.Getenv(("DB_HOST")), os.Getenv(("DB_PORT")), os.Getenv(("DB_NAME"))

	db, err := Connect(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		t.Fatalf("Failed to open a database connection: %v", err)
	}

	// Apparently ping verifies a connection to the database is still alive
	err = db.Ping()
	if err != nil {
		t.Fatalf("Database is unreachable %v", err)
	}

	t.Log("Database connection successful")

}
