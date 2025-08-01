package main

import (
	"TodoApp/Api/controllers"
	"TodoApp/Api/routes"
	"TodoApp/Application/services"
	"TodoApp/infrastructure/database"
	"TodoApp/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Seed()

	r := gin.Default()

	userRepo := &repository.UserRepository{}
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.RegisterRoutes(r, userController)
	r.Run(":8080")

}
