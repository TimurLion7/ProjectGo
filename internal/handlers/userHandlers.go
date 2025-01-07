package handlers

import (
	"context"
	"myproject/internal/models"
	"myproject/internal/userService"
	"myproject/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// DeleteUsersId implements users.StrictServerInterface.
func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := u.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUsersId204Response{}
	return response, nil
}

// GetUsers implements users.StrictServerInterface.
func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (u *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id

	updates := make(map[string]interface{})

	if request.Body.Email != nil {
		updates["email"] = *request.Body.Email
	}

	if request.Body.Password != nil {
		updates["password"] = *request.Body.Password
	}

	updatedUser, err := u.Service.PatchUserByID(id, updates)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем пользователя
	userToCreate := models.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdTask, err := u.Service.PostUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:       &createdTask.ID,
		Email:    &createdTask.Email,
		Password: &createdTask.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}
