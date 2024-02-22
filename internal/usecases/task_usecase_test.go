package usecase

import (
	"testing"

	entities "github.com/GabrielAp542/goTask/internal/entities"
	"github.com/GabrielAp542/goTask/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	//**test connection parameters currently by string until dev container works
	dsn := "host=172.18.0.2 user=postgres password=1234 dbname=test_tasksDB port=5432"
	// Postgres conection by getting env variables
	//open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//detect any error
	if err != nil {
		panic("Failed to connect to database - closing api")
	} else {
		// Migrar esquemas
		db.AutoMigrate(&entities.Task{})
		db.AutoMigrate(&entities.Users{})
	}
	return db
}

// MockDB representa una base de datos ficticia

// CreateTask simula la creación de una tarea en la base de datos
// TestCreateTask testea la función CreateTask
func TestCreateTasks(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	// Creamos una instancia de MockDB

	// Creamos una tarea válida
	Task := &entities.Task{
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}

	// Probamos crear una tarea válida
	err := taskUseCase.CreateTask(Task)
	if err != nil {
		t.Errorf("Se esperaba error nulo para una tarea válida, se recibió: %v", err)
	}

	// Probamos crear una tarea inválida (sin nombre)
	Tasknot := &entities.Task{Task_name: ""}
	err2 := taskUseCase.CreateTask(Tasknot)
	if err2 == nil {
		t.Error("Se esperaba un error para una tarea sin nombre, pero no se recibió ninguno")
	}
}

func TestGetTasks(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	// Creamos una instancia de MockDB

	// Creamos una tarea válida
	//Task := &entities.Task{Task_name: "uwu"}

	// Probamos crear una tarea válida
	tasks, err := taskUseCase.GetTasks()
	if tasks == nil {
		t.Errorf("No se encontro struct de resultado")
	}
	if err != nil {
		t.Errorf("Se esperaba error nulo para una tarea válida, se recibió: %v", err)
	}

}

func TestGetTask(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}
	// Creamos una instancia de MockDB
	taskUseCase.CreateTask(taskCr)
	// Creamos una tarea válida
	//Task := &entities.Task{Task_name: "uwu"}

	// Probamos crear una tarea válida

	tasks, err := taskUseCase.GetTask(uint(IdTask))
	if tasks.TaskId != IdTask {
		t.Errorf("error al obtener los tasks, id_task = %d, recived = %d", IdTask, tasks.TaskId)
	}
	if err != nil {
		t.Errorf("Se esperaba error nulo para una tarea válida, se recibió: %v", err)
	}

}

func TestUpdateTask(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}
	// Creamos una instancia de MockDB
	taskUseCase.CreateTask(taskCr)
	// Creamos una tarea válida
	//Task := &entities.Task{Task_name: "uwu"}
	task := &entities.Task{
		TaskId: IdTask,
	}
	// Probamos crear una tarea válida

	err := taskUseCase.UpdateTask(task)
	if err != nil {
		t.Errorf("Error al actualizar, %s", err)
	}
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	IdTask := 1
	taskCr := &entities.Task{
		TaskId:    IdTask,
		Task_name: "uwu",
		Completed: true,
		Id_User:   nil,
	}
	// Creamos una instancia de MockDB
	taskUseCase.CreateTask(taskCr)
	// Creamos una tarea válida
	// Probamos crear una tarea válida
	err := taskUseCase.DeleteTask(uint(IdTask))
	if err != nil {
		t.Errorf("Error al eliminar, %s", err)
	}
}
