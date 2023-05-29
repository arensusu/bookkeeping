package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbname   = "website"
	sslmode  = "disable"
	timezone = "Asia/Taipei"
)

var err error
var Database *gorm.DB

func Connect() {
	var DB_HOST = os.Getenv("DATABASE_HOST")
	var DB_PORT = os.Getenv("DATABASE_PORT")
	var DB_USER = os.Getenv("DATABASE_USER")
	var DB_PASSWORD = os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, dbname, sslmode, timezone)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
		_ = Database
	}
}
