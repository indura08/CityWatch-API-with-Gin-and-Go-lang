package User

import "citywatch/internal/enums"

type RegisterDto struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	District  enums.District
	Province  enums.Province
}
