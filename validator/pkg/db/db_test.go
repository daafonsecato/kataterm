package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenDB(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=your_username password=your_password dbname=your_database_name sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	t.Log("Success")
}
