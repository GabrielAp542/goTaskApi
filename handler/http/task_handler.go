// delivery/http/task_handler.go
package http

import (
	"net/http"
	"strconv"

	entities "github.com/GabrielAp542/goTask/internal/1entities"
	usecase "github.com/GabrielAp542/goTask/internal/2usecases"
	"github.com/gin-gonic/gin"
)

// struct TaskHandler with taskUseCase field which is defineted with the type TaskUseCase
type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

// func that creats an instance for TaskHandler, recives the abstraccion from UseCases
// so to acess and returns it to be used in the package
func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

// function that implements the struck TaskHandler and recives as a parameter a gin.Context
// to process the errors
func (h *TaskHandler) CreateTask(c *gin.Context) {
	//entities
	var task entities.Task
	//a JSON instance is filled with the struck fiels
	//if an error is detected, it stops the code
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//calls the database to execute the operation
	//if an error is detected, the code stops with its respective error
	if err := h.taskUseCase.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//confirmation of the success of the operation
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := h.taskUseCase.GetTask(uint(taskID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting the values"})
		return
	}
	tasks, err := h.taskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// función actualizar task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask entities.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.TaskId = int(taskID)
	if err := h.taskUseCase.UpdateTask(&updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.taskUseCase.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting task"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
