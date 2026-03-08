package service

import (
	"citywatch/internal/DTO/Incident"
	"citywatch/internal/enums"
	"citywatch/internal/models"
	"citywatch/internal/repository"
	"errors"
)

type IncidentService struct {
	incidentRepository *repository.IncidentRepository
	userRepository     *repository.UserRepository
}

func NewIncidentService(i *repository.IncidentRepository, u *repository.UserRepository) IncidentService {
	return IncidentService{incidentRepository: i, userRepository: u}
}

func (i *IncidentService) CreateIncident(incidentDto *Incident.IncidentDto, imagePath string) error {
	existingUser, err := i.userRepository.GetUserByUserId(incidentDto.ReportedByUserId)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("The reporting user cannot be found , please check again")
	}

	newIncident := models.Incident{Description: incidentDto.Description,
		IncidentCategory:   incidentDto.Category,
		IncidentStatus:     enums.ACTIVE,
		Longitude:          incidentDto.Longitude,
		Latitude:           incidentDto.Latitude,
		Address:            incidentDto.Address,
		IslocationVerified: incidentDto.IsLocationVerified,
		ImageUrl:           imagePath,
		ReportedByUserId:   incidentDto.ReportedByUserId,
		AssignedtoUserId:   incidentDto.AssignedToUserId,
	}
	//image save wena eka hdnne controller ekn, service ekn newei

	repoErr := i.incidentRepository.CreateIncident(&newIncident)
	if repoErr != nil {
		return repoErr
	}

	return nil
}

func (i *IncidentService) DeleteIncidentById(incidentId int) error {
	exisitingIncident, err := i.incidentRepository.GetIncidentById(incidentId)

	if err != nil {
		return err
	}

	if exisitingIncident == nil {
		return errors.New("Incident cannot be found with the privided id")
	}

	result := i.incidentRepository.DeleteIncidentByID(incidentId)
	if result != nil {
		return result
	}

	return nil
}

func (i *IncidentService) GetAllIncidents() ([]models.Incident, error) {

	incidentList, err := i.incidentRepository.GetIncidents()
	if err != nil {
		return nil, err
	}

	return incidentList, nil
}

func (i *IncidentService) AssignWorkerToIncident(incidentId int, userId int) error {
	currentIncident, err1 := i.incidentRepository.GetIncidentById(incidentId)
	if err1 != nil {
		return err1
	}

	if currentIncident == nil {
		return errors.New("Cannot find a incident with the provided ID")
	}

	currentUser, err2 := i.userRepository.GetUserByUserId(userId)
	if err2 != nil {
		return err2
	}

	if currentUser == nil {
		return errors.New("Cannot find a user with provided user ID")
	}

	if currentUser.Role == 2 {
		return errors.New("Selected user is not a worker user!")
	}

	currentIncident.AssignedtoUserId = currentUser.ID

	err3 := i.incidentRepository.UpdateIncident(currentIncident)
	if err3 != nil {
		return err3
	}

	return nil
}

func (i *IncidentService) UpdateIncidentStatus(incidentId int, incidentStatus enums.IncidentStatus) error {
	currentIncident, err1 := i.incidentRepository.GetIncidentById(incidentId)
	if err1 != nil {
		return err1
	}

	if currentIncident == nil {
		return errors.New("Cannot find a incident with the provided ID")
	}

	currentIncident.IncidentStatus = incidentStatus
	err2 := i.incidentRepository.UpdateIncident(currentIncident)
	if err2 != nil {
		return err2
	}

	return nil
}

func (i *IncidentService) UpdateIncident(incidentId int, updatedIncident *Incident.IncidentDto) error {
	exisitingIncident, err1 := i.incidentRepository.GetIncidentById(incidentId)
	if err1 != nil {
		return err1
	}

	if exisitingIncident == nil {
		return errors.New("Could not find a incident with provided ID")
	}

	exisitingIncident.Address = updatedIncident.Address
	exisitingIncident.Description = updatedIncident.Description
	exisitingIncident.Latitude = updatedIncident.Latitude
	exisitingIncident.Longitude = updatedIncident.Longitude
	exisitingIncident.IncidentCategory = updatedIncident.Category
	exisitingIncident.IslocationVerified = updatedIncident.IsLocationVerified
	//anything else cannot be updated due to business logic

	err2 := i.incidentRepository.UpdateIncident(exisitingIncident)
	if err2 != nil {
		return err2
	}

	return nil
}
