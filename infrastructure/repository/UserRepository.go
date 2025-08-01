package repository

import (
	"TodoApp/domain/models"
	"TodoApp/infrastructure/database"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}

type UserRepository struct{}

func (u *UserRepository) Create(user *models.User) error {
	if err := database.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User

	result := database.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
