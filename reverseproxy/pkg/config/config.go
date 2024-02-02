package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	// Add other configuration fields as needed
}

func LoadConfig() (*Config, error) {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     5432,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
		// Load other configurations as needed
	}, nil
}
