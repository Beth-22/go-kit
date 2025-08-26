package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all environment variables
type Config struct {
	DatabaseURL string
	JWTSecret   string
	HasuraAdmin string
}

// LoadConfig loads environment variables from .env file (if exists)
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	hasuraAdmin := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")

	if dbURL == "" || jwtSecret == "" {
		log.Fatal("Missing required environment variables DATABASE_URL or JWT_SECRET")
	}

	return &Config{
		DatabaseURL: dbURL,
		JWTSecret:   jwtSecret,
		HasuraAdmin: hasuraAdmin,
	}
}
