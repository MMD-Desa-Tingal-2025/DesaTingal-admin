package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgreSQL connection
func NewPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Check if database is reachable
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
