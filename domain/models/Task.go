package models

import (
	"TodoApp/Application/marshalJSON"
)

type Task struct {
	ID          uint             `json:"id" gorm:"primaryKey"`
	ProjectId   uint             `json:"projectId" gorm:"index"`
	SectionId   *uint            `json:"sectionId" gorm:"index"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Date        marshalJSON.Date `json:"date"`
	Priority    string           `json:"priority"`
	
	Completed bool `json:"completed"`

	Section *Section `json:"section" gorm:"foreignKey:SectionId"`
	Project Project  `json:"project" gorm:"foreignKey:ProjectId"`
}
