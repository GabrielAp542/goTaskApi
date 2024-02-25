package usecase

import (
	"testing"

	"github.com/GabrielAp542/goTask/cmd/database"
	entities "github.com/GabrielAp542/goTask/internal/entities"
	"github.com/GabrielAp542/goTask/internal/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// MockDB representa una base de datos ficticia

// CreateTask simula la creación de una tarea en la base de datos
// TestCreateTask testea la función CreateTask

// connection variables
var db *gorm.DB
var errDB error

func init() {
	db, errDB = database.TestingDB(false)
}

func TestCreateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	task := &entities.Task{
		Task_name: "task",
		Completed: true,
		Id_User:   nil,
	}
	err := taskUseCase.CreateTask(task)
	assert.NoError(t, err)

	taskNoName := &entities.Task{
		Task_name: "",
		Completed: true,
		Id_User:   nil,
	}
	errNoName := taskUseCase.CreateTask(taskNoName)
	assert.Error(t, errNoName)

}

func TestGetTasks(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)

	// Probamos crear una tarea válida
	tasks, errR := taskUseCase.GetTasks()
	assert.NotNil(t, tasks)
	assert.NoError(t, errR)
}

func TestGetTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "task",
		Completed: true,
		Id_User:   nil,
	}
	taskUseCase.CreateTask(taskCr)
	tasks, err := taskUseCase.GetTask(uint(IdTask))
	assert.Equal(t, tasks.TaskId, IdTask)
	assert.NoError(t, err)

}

func TestUpdateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}
	taskUseCase.CreateTask(taskCr)
	task := &entities.Task{
		TaskId: IdTask,
	}
	err := taskUseCase.UpdateTask(task)
	assert.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	assert.NoError(t, errDB)

	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)

	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}
	taskUseCase.CreateTask(taskCr)
	err := taskUseCase.DeleteTask(uint(IdTask))
	assert.NoError(t, err)
}
