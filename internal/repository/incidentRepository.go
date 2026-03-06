package repository

import (
	"citywatch/internal/models"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (i *IncidentRepository) CreateIncident(incident *models.Incident) error {

	result := i.db.Create(&incident)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (i *IncidentRepository) DeleteIncidentByID(incidentId int) error {
	//lets find the incident first
	var currentIncident *models.Incident

	result := i.db.First(&currentIncident, incidentId)
	if result.Error != nil {
		return result.Error
	}

	//then delete it
	deleteResult := i.db.Delete(&currentIncident)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}

func (i *IncidentRepository) GetIncidentById() ([]models.Incident, error) {
	var incidents []models.Incident
	result := i.db.Find(&incidents)
	//methanid wenne gorm ekn pass krna variable eke type ekt adlwa struct eke thiyna name eka lowecase krla thami table name ek hdaganne, meka gorm wal thiyna honda wishesa deyk

	if result.Error != nil {
		return nil, result.Error
	}

	return incidents, nil
}
