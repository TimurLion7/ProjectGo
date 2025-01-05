package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	PostUser(user User) (User, error)
	PatchUserByID(id int, patchs map[string]interface{}) (User, error)
	DeleteUserByID(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	var Users []User

	err := r.db.Order("Created_at ASC").Find(&Users).Error

	return Users, err
}

func (r *userRepository) PostUser(user User) (User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

func (r *userRepository) PatchUserByID(id int, patchs map[string]interface{}) (User, error) {
	res := r.db.Model(&User{}).Where("id = ?", id).Updates(patchs)
	if res.Error != nil {
		return User{}, res.Error
	}

	var UpdateUser User

	if err := r.db.Find(&UpdateUser, id).Error; err != nil {
		return User{}, err
	}
	return UpdateUser, nil
}

func (r *userRepository) DeleteUserByID(id int) error {
	var user User
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
