package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"` // Внешний ключ на User
	User   User   `json:"user"`    // Связь с пользователем
}
