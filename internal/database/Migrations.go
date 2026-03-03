package database

import (
	"citywatch/internal/database"
	"citywatch/internal/models"
)

func init() {
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{}, &models.Incident{})
}

//need to put this migration code inside main.go file to avoid the circular package depandacy error
