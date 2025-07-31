package routes

import (
	"TodoApp/Api/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.PUT("/", controllers.UpdateTask)
		taskGroup.DELETE("/", controllers.DeleteTask)
	}

	sectionGroup := r.Group("/sections")
	{
		sectionGroup.GET("/", controllers.GetSections)
		sectionGroup.POST("/", controllers.CreateSection)
		sectionGroup.DELETE("/", controllers.DeleteSection)
		sectionGroup.PUT("/", controllers.UpdateSection)
	}
}
