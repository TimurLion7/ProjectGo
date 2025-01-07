package taskService

import (
	"myproject/internal/models"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id int, updates map[string]interface{}) (models.Task, error) {
	return s.repo.UpdateTaskByID(id, updates)
}

func (s *TaskService) DeleteTaskByID(id int) error {
	return s.repo.DeleteTaskByID(id)
}
