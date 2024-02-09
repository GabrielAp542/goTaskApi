// repository/task_repository.go
package repository

import (
	entities "github.com/GabrielAp542/goTask/internal/1entities"
	"gorm.io/gorm"
)

// struct that defines the gorm package so it can be used on the code
type TaskRepository struct {
	db *gorm.DB
}

// func that recives the conection data and returns it
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *entities.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	return tasks, r.db.Find(&tasks).Error
}

func (r *TaskRepository) UpdateTask(task *entities.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&entities.Task{}, id).Error
}
