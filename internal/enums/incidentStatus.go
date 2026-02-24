package enums

type IncidentStatus int

const (
	ACTIVE   IncidentStatus = 0
	INACTIVE IncidentStatus = 1
	ONHOLD   IncidentStatus = 2
	REJECTED IncidentStatus = 3
	DONE     IncidentStatus = 4
)
