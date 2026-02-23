package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var error error

	dsn := "database-connection-string"
	//methna database connection string eka ganna .env ekn = os.Getenv("DATABASE_URL")

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Couldnt Connect to database")
	}

	DB = db
}
