package userService

import (
	"myproject/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	PostUser(user models.User) (models.User, error)
	PatchUserByID(id int, patchs map[string]interface{}) (models.User, error)
	DeleteUserByID(id int) error
	GetTasksForUser(userID uint) ([]models.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetTasksForUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var Users []models.User

	err := r.db.Order("Created_at ASC").Find(&Users).Error

	return Users, err
}

func (r *userRepository) PostUser(user models.User) (models.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return user, nil
}

func (r *userRepository) PatchUserByID(id int, patchs map[string]interface{}) (models.User, error) {
	res := r.db.Model(&models.User{}).Where("id = ?", id).Updates(patchs)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	var UpdateUser models.User

	if err := r.db.Find(&UpdateUser, id).Error; err != nil {
		return models.User{}, err
	}
	return UpdateUser, nil
}

func (r *userRepository) DeleteUserByID(id int) error {
	var user models.User
	res := r.db.First(&user, id)
	if res.Error != nil {
		return res.Error // Вернуть ошибку, если задача не найдена
	}
	DeleteRes := r.db.Delete(&user)
	if DeleteRes.Error != nil {
		return DeleteRes.Error
	}
	return nil
}
