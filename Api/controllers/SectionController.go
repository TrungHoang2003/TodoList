package controllers

import (
	"TodoApp/domain/models"
	"TodoApp/infrastructure/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSections(c *gin.Context) {
	var sections []models.Section
	if err := database.Db.Find(&sections).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": sections})
}

func CreateSection(c *gin.Context) {
	var section models.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	database.Db.Create(&section)
	c.JSON(http.StatusOK, gin.H{"data": section})
}

func DeleteSection(c *gin.Context) {
	var section models.Section
	if err := database.Db.Delete(&section).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message ": "Section deleted successfully"})
}

func UpdateSection(c *gin.Context) {
	var section models.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	database.Db.Save(&section)
	c.JSON(http.StatusOK, gin.H{"message": "Section updated successfully", "data": section})
}
