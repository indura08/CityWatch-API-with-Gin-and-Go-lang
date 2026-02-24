package service

import "citywatch/internal/repository"

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

//meke methods tika implement krnna elaga dwse meka krddi
