package repositories

import (
	"testing"

	entities "github.com/GabrielAp542/goTaskApi/internal/entities"
	"github.com/GabrielAp542/goTaskApi/setup/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// connection variables
var db, dbF *gorm.DB
var errDB, errDBF error

func init() {
	db, errDB = database.TestingDB(false)
	dbF, errDBF = database.TestingDB(true)
}

func TestCreateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := NewTaskRepository(db)
	Task := &entities.Task{
		Task_name: "prueba",
		Completed: true,
	}
	errR := taskRepo.CreateTask(Task)
	assert.NoError(t, errR)
}

func TestGetTasks(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := NewTaskRepository(db)
	tasks, errR := taskRepo.GetTasks()
	assert.NoError(t, errR)
	assert.NotNil(t, tasks)
}

func TestGetTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := NewTaskRepository(db)
	Task := &entities.Task{
		TaskId:    1,
		Task_name: "uwu",
		Completed: true,
	}
	taskRepo.CreateTask(Task)
	IdTask := 1
	tasks, errR := taskRepo.GetTask(uint(IdTask))
	assert.Equal(t, tasks.TaskId, IdTask)
	assert.NoError(t, errR)
}

func TestUpdateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := NewTaskRepository(db)
	TaskCr := &entities.Task{
		TaskId:    1,
		Task_name: "asd",
		Completed: true,
	}
	//--update succesfull
	taskRepo.CreateTask(TaskCr)
	IdTask := 1
	Task := &entities.Task{
		Task_name: "asd_update",
		Completed: false,
		Id_User:   nil,
		TaskId:    IdTask,
	}
	errR := taskRepo.UpdateTask(Task)
	assert.NoError(t, errR)
}

func TestDeleteTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := NewTaskRepository(db)
	IdTask := 1
	errR := taskRepo.DeleteTask(uint(IdTask))
	assert.NoError(t, errR)
}
