package taskService

import (
	"myproject/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task models.Task) (models.Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]models.Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id int, updates map[string]interface{}) (models.Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	// Проверка, что user_id существует
	var user []models.User
	if err := r.db.First(&user, task.UserID).Error; err != nil {
		return models.Task{}, err
	}
	// Создание задачи
	result := r.db.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	//err := r.db.Find(&tasks).Error
	err := r.db.Order("created_at ASC").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id int, updates map[string]interface{}) (models.Task, error) {

	result := r.db.Model(&models.Task{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return models.Task{}, result.Error
	}

	var updatedTask models.Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return models.Task{}, err
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id int) error {
	var task models.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return result.Error // Вернуть ошибку, если задача не найдена
	}
	deleteResult := r.db.Delete(&task)
	if deleteResult.Error != nil {
		return deleteResult.Error // Вернуть ошибку, если удаление не удалось
	}

	return nil // Успешное выполнение

}
