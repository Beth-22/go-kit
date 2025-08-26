package main

import (
	"moviedb-hasura/internal/config"
	"moviedb-hasura/internal/db"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to Postgres
	db.Connect(cfg)

	// Test query
	var now string
	err := db.DB.QueryRow(db.DB.Context(), "SELECT NOW()").Scan(&now)
	if err != nil {
		panic(err)
	}
	println("Database connection test successful! Current time:", now)
}
