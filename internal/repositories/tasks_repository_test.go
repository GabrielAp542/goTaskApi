package repositories

import (
	"fmt"
	"testing"

	entities "github.com/GabrielAp542/goTask/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(host string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=postgres password=1234 dbname=test_tasksDB port=5432", host)
	// Postgres conection by getting env variables
	/*sn := fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s  port=%s",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("TEST_DB_NAME"),
	os.Getenv("DB_PORT"))*/
	//open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//detect any error
	if err != nil {
		//panic("Failed to connect to database - closing api")
	}
	// Migrar esquemas
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Users{})
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTestDB("172.18.0.2")
	taskRepo := NewTaskRepository(db)
	Task := &entities.Task{
		Task_name: "uwu",
		Completed: true,
	}
	err := taskRepo.CreateTask(Task)
	if err != nil {
		t.Errorf("error detectado")
	}

	dbf := setupTestDB("uwu")
	taskRepof := NewTaskRepository(dbf)
	errf := taskRepof.CreateTask(Task)
	if errf == nil {
		t.Errorf("error detectado")
	}
}

func TestGetTasks(t *testing.T) {
	db := setupTestDB("172.18.0.2")
	taskRepo := NewTaskRepository(db)
	tasks, err := taskRepo.GetTasks()
	if tasks == nil {
		t.Errorf("no hay tasks")
	}
	if err != nil {
		t.Errorf("error detectado")
	}
	/*
		dbf := setupTestDB("uwu")
		taskRepof := NewTaskRepository(dbf)
		errf := taskRepof.CreateTask(Task)
		if errf == nil {
			t.Errorf("error detectado")
		}*/
}

func TestGetTask(t *testing.T) {
	db := setupTestDB("172.18.0.2")
	taskRepo := NewTaskRepository(db)
	Task := &entities.Task{
		TaskId:    1,
		Task_name: "uwu",
		Completed: true,
	}
	taskRepo.CreateTask(Task)
	IdTask := 1
	tasks, err := taskRepo.GetTask(uint(IdTask))
	if tasks.TaskId != IdTask {
		t.Errorf("error al obtener los tasks, id_task = %d, recived = %d", IdTask, tasks.TaskId)
	}
	if err != nil {
		t.Errorf("error detectado, %s", err)
	}
	/*
		dbf := setupTestDB("uwu")
		taskRepof := NewTaskRepository(dbf)
		errf := taskRepof.CreateTask(Task)
		if errf == nil {
			t.Errorf("error detectado")
		}*/
}

func TestUpdateTask(t *testing.T) {
	db := setupTestDB("172.18.0.2")
	taskRepo := NewTaskRepository(db)
	TaskCr := &entities.Task{
		TaskId:    1,
		Task_name: "uwu",
		Completed: true,
	}
	taskRepo.CreateTask(TaskCr)
	IdTask := 1
	Task := &entities.Task{
		Task_name: "uwu_update",
		Completed: true,
		Id_User:   nil,
		TaskId:    IdTask,
	}
	err := taskRepo.UpdateTask(Task)
	if err != nil {
		t.Errorf("error al actualzar, %s", err)
	}
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB("172.18.0.2")
	taskRepo := NewTaskRepository(db)
	IdTask := 1
	err := taskRepo.DeleteTask(uint(IdTask))
	if err != nil {
		t.Errorf("error detectado al eliminar")
	}
	/*
		dbf := setupTestDB("uwu")
		taskRepof := NewTaskRepository(dbf)
		errf := taskRepof.CreateTask(Task)
		if errf == nil {
			t.Errorf("error detectado")
		}*/
}