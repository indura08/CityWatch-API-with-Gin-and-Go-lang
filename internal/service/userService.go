package service

import (
	auth "citywatch/internal/Dto/Auth"
	user "citywatch/internal/Dto/User"
	"citywatch/internal/models"
	"citywatch/internal/repository"
	"citywatch/internal/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (u *UserService) Login(loginDto *auth.LoginDto) (string, error) {
	existingUser, err1 := u.userRepository.FindUserByEmail(loginDto.Email)
	if err1 != nil {
		return "", err1
	}

	if existingUser == nil {
		return "", fmt.Errorf("Wrong credentials")
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(loginDto.Password))
	if err2 != nil {
		return "", fmt.Errorf("Wrong Credentials")
	}

	token := utils.JwtTokenGenerator(loginDto.Email, existingUser.Role, existingUser.ID)
	if token == "" {
		return "", fmt.Errorf("Something went wrong please try again later")
	}

	return token, nil

}

// meke methods tika implement krnna elaga dwse meka krddi
func (u *UserService) Register(registerDto *user.RegisterDto) error {
	existingUser, err := u.userRepository.FindUserByEmail(registerDto.Email)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return fmt.Errorf("User already exists")
	}

	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err1 != nil {
		return err1
	}

	//methna user ge role ek hdna mechanism ek ghnna

	newUser := models.User{FirstName: registerDto.FirstName, LastName: registerDto.LastName, Email: registerDto.Email,
		District: registerDto.District, Province: registerDto.Province,
		PasswordHash: string(hashedPassword)}

	return u.userRepository.CreateNewUser(&newUser)

}
