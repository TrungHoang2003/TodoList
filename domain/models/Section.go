package models

type Section struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProjectId uint   `json:"projectId" gorm:"index"`
	Name      string `json:"name"`

	Project Project `json:"project" gorm:"foreignKey:ProjectId"`
}
