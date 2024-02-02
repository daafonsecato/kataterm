package database

import (
	"database/sql"
	"fmt"

	"github.com/daafonsecato/kataterm-reverseproxy/pkg/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
		return nil, err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

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
