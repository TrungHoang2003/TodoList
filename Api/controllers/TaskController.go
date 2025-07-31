package controllers

import (
	"TodoApp/domain/models"
	"TodoApp/infrastructure/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.Db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message ": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Db.Save(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Db.Create(&task)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := database.Db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
