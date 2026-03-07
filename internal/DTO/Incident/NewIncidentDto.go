package Incident

import (
	"citywatch/internal/enums"
	"mime/multipart"
)

type IncidentDto struct {
	Description string
	Category    enums.IncidentCategory

	Latitude  float32
	Longitude float32

	Address            string
	IsLocationVerified bool

	Image *multipart.FileHeader

	ReportedByUserId int
	AssignedToUserId int
}
