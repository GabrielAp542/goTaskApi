package usecase

import (
	"testing"

	entities "github.com/GabrielAp542/goTask/internal/entities"
	"github.com/GabrielAp542/goTask/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "host=172.18.0.2 user=postgres password=1234 dbname=test_tasksDB port=5432"
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
func TestCreateTask(t *testing.T) {
	db := setupTestDB()
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := NewTaskUseCase(taskRepository)
	// Creamos una instancia de MockDB

	// Creamos una tarea válida
	Task := &entities.Task{Task_name: "uwu"}

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
