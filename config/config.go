package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	ServerPort string
}

func Load() (*Config, error) {
	_ = godotenv.Load() // optional .env

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return &Config{DBUrl: dsn, ServerPort: ":" + port}, nil
}
