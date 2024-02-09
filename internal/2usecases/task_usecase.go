// usecase/task_usecase.go
package usecase

import (
	entities "github.com/GabrielAp542/goTask/internal/1entities"
)

// interface with all the methods to be implemented by the use cases
type TaskRepository interface {
	CreateTask(task *entities.Task) error
	GetTasks() ([]entities.Task, error)
	UpdateTask(task *entities.Task) error
	DeleteTask(id uint) error
}

// struck type TaskUseCase with a field corresponding to a dependency injection
type TaskUseCase struct {
	taskRepository TaskRepository
}

// function that takes the struct as a parameter and returns a pinter of that one
func NewTaskUseCase(taskRepository TaskRepository) *TaskUseCase {
	//returns the value of the repository sent before
	return &TaskUseCase{taskRepository: taskRepository}
}

// business logic to create a task
func (uc *TaskUseCase) CreateTask(task *entities.Task) error {

	return uc.taskRepository.CreateTask(task)
}

// business logic to get a task
func (uc *TaskUseCase) GetTasks() ([]entities.Task, error) {
	return uc.taskRepository.GetTasks()
}

// business logic to update a task, uses the struct taskusecase as uc, recives entities as parameter

func (uc *TaskUseCase) UpdateTask(task *entities.Task) error {
	//return the call from the repository sending entities as a parameter
	return uc.taskRepository.UpdateTask(task)
}

// business logic to delete a task
func (uc *TaskUseCase) DeleteTask(id uint) error {
	return uc.taskRepository.DeleteTask(id)
}
