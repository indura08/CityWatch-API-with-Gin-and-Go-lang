package models

import (
	"citywatch/internal/enums"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserID       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	Role         enums.Role

	District enums.District
	Province enums.Province

	//gorm.Model eke created at updated at thiyna hinda awlk nha

	//heta meke literal strings tika hdla ithuru tika continue krn yanna blnna
}
