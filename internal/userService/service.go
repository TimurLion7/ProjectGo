package userService

type UserService struct {
	UserRepo UserRepository
}

func NewUserService(UserRepo UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.UserRepo.GetUsers()
}

func (s *UserService) PostUser(user User) (User, error) {
	return s.UserRepo.PostUser(user)
}

func (s *UserService) PatchUserByID(id int, patchs map[string]interface{}) (User, error) {
	return s.UserRepo.PatchUserByID(id, patchs)
}

func (s *UserService) DeleteUserByID(id int) error {
	return s.UserRepo.DeleteUserByID(id)
}
