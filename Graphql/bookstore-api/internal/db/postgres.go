package db

import (
    "os"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
	"github.com/joho/godotenv"
)


var Pool *pgxpool.Pool

func Connect() (error) {
    // get database URL from environment variables
	err := godotenv.Load()
	if err!=nil{
		return fmt.Errorf("error loading .env file: %w", err)
	}
		dbURL := os.Getenv("DB_URL")


    // open a connection to Postgres
   Pool, err = pgxpool.New(context.Background(), dbURL)

   if err != nil{
	return fmt.Errorf("unable to create a connection pool: %w", err)
   }
    // ping to check connection
	err = Pool.Ping(context.Background())

	if err != nil{
		return fmt.Errorf("unable to connect to db: %w", err)
	}

    // return db and error
	fmt.Println("connected to postgres db succesfully!")
	return nil
}
