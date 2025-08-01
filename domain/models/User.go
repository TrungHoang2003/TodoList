package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;"`
	Username string `json:"name" gorm:"not null;unique;size:255"`
	Password string `json:"password" gorm:"not null"`
}
