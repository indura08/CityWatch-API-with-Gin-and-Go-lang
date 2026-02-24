package models

import (
	"citywatch/internal/enums"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           int        `json:"id" gorm:"primaryKey"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	Email        string     `json:"email" gorm:"not null;unique"`
	PasswordHash string     `json:"password" gorm:"not null"`
	Role         enums.Role `json:"role" gorm:"not null"`

	District enums.District `json:"district"`
	Province enums.Province `json:"province"`

	//gorm.Model eke created at updated at thiyna hinda awlk nha

	//heta meke literal strings tika hdla ithuru tika continue krn yanna blnna
}
