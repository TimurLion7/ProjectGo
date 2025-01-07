package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"` // Список задач, связанных с пользователем
}
