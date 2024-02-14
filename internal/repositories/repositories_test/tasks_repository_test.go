package repositoriestest_test

import (
	"log"
	"testing"

	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
	type MockDB struct {
		mock.Mock
	}

// func that recives the conection data and returns it

	func NewTaskRepository(db *gorm.DB) *TaskRepository {
		return &TaskRepository{db: db}
	}

func TestCreateTask_Success(t *testing.T) {

}

	func (r *TaskRepository) CreateTask(task *entities.Task) error {
		return r.db.Create(task).Error
	}
*/
type Database interface {
	Create(*entities.Task) error
	// Agrega otros métodos necesarios aquí
}

type TaskRepository struct {
	db Database
}

// Mock de la base de datos
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(task *entities.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestCreateTask_Success(t *testing.T) {
	// Crear instancia del mock de la base de datos
	mockDB := new(MockDB)

	// Configurar comportamiento esperado para el método Create en el mock
	mockDB.On("Create", mock.AnythingOfType("*entities.Task")).Return(nil)

	// Crear instancia del repositorio de tareas con el mock de la base de datos
	taskRepo := TaskRepository{db: mockDB}

	id := 0
	ptr := &id
	// Crear una tarea para la prueba
	task := &entities.Task{
		TaskId:    1,
		Task_name: "Task 1",
		Completed: true,
		Id_User:   ptr,
	}

	// Llamar a la función CreateTask del repositorio de tareas
	err := taskRepo.db.Create(task)
	log.Println("Mensaje de debug")
	// Verificar que no se haya producido un error
	assert.NoError(t, err)

	// Verificar que se haya llamado al método Create en la base de datos con la tarea correcta
	mockDB.AssertCalled(t, "Create", task)
}

func TestCreateTask_Fail(t *testing.T) {
	// Crear instancia del mock de la base de datos
	mockDB := new(MockDB)
	// Configurar comportamiento esperado para el método Create en el mock
	mockDB.On("Create", mock.AnythingOfType("*entities.Task")).Return(nil)

	// Crear instancia del repositorio de tareas con el mock de la base de datos
	taskRepo := TaskRepository{db: mockDB}

	// Crear una tarea para la prueba
	task := &entities.Task{}

	// Llamar a la función CreateTask del repositorio de tareas
	err := taskRepo.db.Create(task)
	log.Println("Mensaje de debug")
	// Verificar que no se haya producido un error
	assert.NoError(t, err)

	// Verificar que se haya llamado al método Create en la base de datos con la tarea correcta
	mockDB.AssertCalled(t, "Create", task)
}
