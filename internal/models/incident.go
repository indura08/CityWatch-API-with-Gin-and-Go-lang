package models

import (
	"citywatch/internal/enums"

	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model

	Id                 int                    `json:"id" gorm:"primaryKey"`
	Description        string                 `json:"description" gorm:"not null"`
	IncidentCategory   enums.IncidentCategory `json:"incidentCategory"`
	IncidentStatus     enums.IncidentStatus   `json:"incidentStatus"`
	Longitude          float32                `json:"longitude"`
	Latitude           float32                `json:"latitude"`
	Address            string                 `json:"address"`
	IslocationVerified bool                   `json:"isLocationVerified"`
	ImageUrl           string                 `json:"imageUrl"`

	ReportedByUserId int `json:"reportedByUserID"`
	ReportedByUser   User

	AssignedtoUserId int `json:"assignedToUserId"`
	AssignedToUser   User
}
