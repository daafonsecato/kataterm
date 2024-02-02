package models

import (
	"database/sql"
	"fmt"

	db "github.com/daafonsecato/kataterm-reverseproxy/internal/database"
)

type SessionStore struct {
	db *sql.DB
}

func NewSessionStore(database *sql.DB) *SessionStore {
	db.InitDB()
	db, err := db.InitDB()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to open database: %v", err)
		panic(errorMsg)
	}

	err = db.Ping()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to ping database: %v", err)
		panic(errorMsg)
	}

	if db == nil {
		panic("db is nil")
	}
	return &SessionStore{
		db: db,
	}
}

func (store *SessionStore) Close() {
	store.db.Close()
}
