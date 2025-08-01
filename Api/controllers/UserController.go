package controllers

import (
	"TodoApp/Application/DTOs"
	"TodoApp/Application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) Register(c *gin.Context) {
	var dto DTOs.UserRegisterDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.service.Register(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func (u *UserController) Login(c *gin.Context) {
	var dto DTOs.UserLoginDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := u.service.Login(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User login successfully!"})
}
