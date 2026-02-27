package user

import "citywatch/internal/enums"

type RegisterDto struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      enums.Role
	District  enums.District
	Province  enums.Province
}
