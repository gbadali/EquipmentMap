package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

// GetDBConnection returns the database connection object
func GetDBConnection() (*sql.DB, error) {
	url := `libsql://libsql://equipment-db-gbadali.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTQ5NzEwNzksImlkIjoiM2Y5OTUyNWEtZTRlNS00ZmQ5LWFhM2MtOWM3ODM4ZGM0ODg5In0.A3cTFVy7f_gIs9PSFJcvtOtqtJwTLRLAnrKq2ljzoOw9OVyj6w_V4Vd1MdeY5etVdGWwEWODNJ3LelXCYmUrBw`

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to opening database: %s: %s", url, err)
		os.Exit(1)
	}

	DB = db
	defer db.Close()
	return db, nil
}
