package models

type Project struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	Sections []Section `json:"sections"`
	Tasks    []Task    `json:"tasks"`
}
