package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     string
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	host = os.Getenv("DB_HOST")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Failed to open database: %v", err)
		return nil, err
	}
	fmt.Printf("Connected to DB")

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
