package main

import (
	"TodoApp/Api/routes"
	"TodoApp/infrastructure/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Seed()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")

}
