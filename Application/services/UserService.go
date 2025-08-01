package services

import (
	"TodoApp/Application/DTOs"
	"TodoApp/domain/models"
	"TodoApp/infrastructure/repository"
	"errors"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Login(dto DTOs.UserLoginDto) error {
	var user *models.User
	user, err := u.repo.GetByUsername(dto.Username)

	if err != nil {
		return err
	}

	if user.Password != dto.Password {
		return errors.New("wrong password")
	}
	return nil
}

func (u *UserService) Register(dto DTOs.UserRegisterDto) error {
	if dto.ConfirmPassword != dto.Password {
		return errors.New("confirm password and password are not the same")
	}

	var user = models.User{
		Username: dto.Username,
		Password: dto.Password,
	}

	err := u.repo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}
