package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "192.168.0.16"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "website"
	sslmode  = "disable"
	timezone = "Asia/Taipei"
)

var err error
var Database *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", host, port, user, password, dbname, sslmode, timezone)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
		_ = Database
	}
}
