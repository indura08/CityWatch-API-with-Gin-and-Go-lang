package repository

import (
	"citywatch/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(user *models.User) error {

	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	error := r.db.Where("email = ?", email).First(&user).Error

	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, error
	}

	return &user, nil

}

func (r *UserRepository) GetUserByUserId(userId int) (*models.User, error) {
	var user models.User
	error := r.db.First(&user, userId).Error

	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, error
	}

	return &user, nil
}
