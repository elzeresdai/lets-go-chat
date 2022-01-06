package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sqlx.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := os.Getenv("DATABASE_URL")
	conn, err := sqlx.Open("postgres", env)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil
	}
	fmt.Println("Successfully connected!")
	return conn
}
