package testusecases_test

import (
	"errors"
	"testing"
)

// MockDB representa una base de datos ficticia
type MockDB struct{}

// CreateTask simula la creación de una tarea en la base de datos
func (uc *MockDB) CreateTask(task *Task) error {
	if task.Task_name == "" {
		return errors.New("el nombre debe ser especificado")
	}
	// En un escenario real, aquí iría la lógica para crear la tarea en la base de datos
	return nil
}

// Task representa una tarea
type Task struct {
	Task_name string
}

// TestCreateTask testea la función CreateTask
func TestCreateTask(t *testing.T) {
	// Creamos una instancia de MockDB
	db := MockDB{}

	// Creamos una tarea válida
	validTask := &Task{Task_name: "uwu"}

	// Probamos crear una tarea válida
	err := db.CreateTask(validTask)
	if err != nil {
		t.Errorf("Se esperaba error nulo para una tarea válida, se recibió: %v", err)
	}

	// Probamos crear una tarea inválida (sin nombre)
	invalidTask := &Task{Task_name: ""}
	err = db.CreateTask(invalidTask)
	if err == nil {
		t.Error("Se esperaba un error para una tarea sin nombre, pero no se recibió ninguno")
	}
}
