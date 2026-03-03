package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var error error

	dsn := os.Getenv("DATABASE_URL")

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Couldnt Connect to database")
	}

	DB = db
}
