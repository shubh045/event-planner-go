package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Could not load .env file")
		return
	}
	
	dbpool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	DB = dbpool

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INT
	)
	`
	_, err := DB.Exec(context.Background(), createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create events table.")
	}
}
