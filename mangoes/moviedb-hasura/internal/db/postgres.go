package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"moviedb-hasura/internal/config"
)

var DB *pgxpool.Pool

// Connect initializes the Postgres connection pool
func Connect(cfg *config.Config) {
	var err error
	DB, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	fmt.Println("Connected to Postgres successfully!")
}
