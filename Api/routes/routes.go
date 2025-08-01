package routes

import (
	"TodoApp/Api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userController *controllers.UserController) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/task", controllers.GetTasks)
		taskGroup.POST("/create-section", controllers.CreateTask)
		taskGroup.PUT("/update-section", controllers.UpdateTask)
		taskGroup.DELETE("/delete-section", controllers.DeleteTask)
	}

	sectionGroup := r.Group("/sections")
	{
		sectionGroup.GET("/sections", controllers.GetSections)
		sectionGroup.POST("/create-section", controllers.CreateSection)
		sectionGroup.DELETE("/delete-section", controllers.DeleteSection)
		sectionGroup.PUT("/update-section", controllers.UpdateSection)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
	}
}
