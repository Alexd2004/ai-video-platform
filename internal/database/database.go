package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname string) (*sql.DB, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	return sql.Open("postgres", connString)
}
