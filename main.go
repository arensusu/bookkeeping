package main

import (
	"bookkeeping-backend/api"
	db "bookkeeping-backend/database/sqlc"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

func main() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	conn, err := sql.Open(dbDriver, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	store := db.New(conn)
	server := api.NewServer(store)
	server.Start(":8080")
}
