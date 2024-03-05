// usecase/task_usecase.go
package usecase

import (
	"errors"

	entities "github.com/GabrielAp542/goTask-Api-Gabriel/internal/entities"
)

// interface with all the methods to be implemented by the use cases
type TaskInterface interface {
	CreateTask(task *entities.Task) error
	GetTasks() ([]entities.Task, error)
	GetTask(id uint) (entities.Task, error)
	UpdateTask(task *entities.Task) error
	DeleteTask(id uint) error
}

// struck type TaskUseCase with a field corresponding to a dependency injection
// contains a type TaskRespository field
type TaskUseCase struct {
	taskRepository TaskInterface
}

// function that takes the struct as a parameter and returns a pinter of that one
func NewTaskUseCase(taskRepository TaskInterface) *TaskUseCase {
	//returns the value of the repository sent before
	return &TaskUseCase{taskRepository: taskRepository}
}

// business logic to create a task
func (uc *TaskUseCase) CreateTask(task *entities.Task) error {
	if task.Task_name == "" {
		return errors.New("the name must ve listed")
	}
	return uc.taskRepository.CreateTask(task)
}

// business logic to get a task
func (uc *TaskUseCase) GetTasks() ([]entities.Task, error) {
	return uc.taskRepository.GetTasks()
}

func (uc *TaskUseCase) GetTask(id uint) (entities.Task, error) {
	return uc.taskRepository.GetTask(id)
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
