package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;"`
	Username string `json:"name" gorm:"unique;"`
	Password string `json:"password"`
}
