package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// GetDBConnection returns the database connection object
func GetDBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:equipment.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
