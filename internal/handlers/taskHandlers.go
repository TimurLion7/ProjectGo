package handlers

import (
	"context"
	"myproject/internal/taskService"
	"myproject/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// DeleteTasksId implements tasks.StrictServerInterface.

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId204Response{}
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id

	// Создаем карту обновляемых полей
	updates := make(map[string]interface{})

	// Проверяем, какие поля были переданы
	if request.Body.Task != nil {
		updates["task"] = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updates["is_done"] = *request.Body.IsDone
	}

	// Обновляем задачу
	updatedTask, err := h.Service.UpdateTaskByID(id, updates)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

/*
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id

	// Теперь предполагаем, что request.Body - это структура с полями Task и IsDone
	update := request.Body // Это структура

	// Создаем объект задачи для обновления
	updTask := taskService.Task{
		Task:   *update.Task,   // предполагается, что Task - это указатель
		IsDone: *update.IsDone, // предполагается, что IsDone - это указатель
	}

	// Обращаемся к сервису для обновления задачи по ID
	updatedTask, err := h.Service.UpdateTaskByID(id, updTask)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}


}*/
