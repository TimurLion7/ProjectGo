package userService

import (
	"myproject/internal/models"
)

type UserService struct {
	UserRepo UserRepository
}

func NewUserService(UserRepo UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.GetUsers()
}

func (s *UserService) PostUser(user models.User) (models.User, error) {
	return s.UserRepo.PostUser(user)
}

func (s *UserService) PatchUserByID(id int, patchs map[string]interface{}) (models.User, error) {
	return s.UserRepo.PatchUserByID(id, patchs)
}

func (s *UserService) DeleteUserByID(id int) error {
	return s.UserRepo.DeleteUserByID(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]models.Task, error) {
	return s.UserRepo.GetTasksForUser(userID)
}
