package database

import (
	"TodoApp/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

const dsn = "root:root@tcp(127.0.0.1:3306)/tasks?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database " + err.Error())
	}

	Db = db

	err = db.AutoMigrate(&models.Task{}, &models.Section{},
		&models.Project{}, &models.User{})
	if err != nil {
		panic("failed to migrate database " + err.Error())
		return
	}
}

func Seed() {
	var count int64
	Db.Model(&models.Project{}).Where("name = ?", "Inbox").Count(&count)

	if count == 0 {
		project := models.Project{
			Name: "Inbox",
		}
		Db.Create(&project)
	}
}
