// delivery/http/task_handler.go
package http

import (
	"net/http"
	"strconv"

	tasks_request "github.com/GabrielAp542/goTask/handler/presenters/requests"
	tasks_response "github.com/GabrielAp542/goTask/handler/presenters/responses"
	usecase "github.com/GabrielAp542/goTask/internal/usecases"

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

// function that implements the struck TaskHandler and recives as a parameter as gin.Context
// to process the errors
func (h *TaskHandler) CreateTask(c *gin.Context) {

	fRequestTask, err := tasks_request.FormatRequestPostandPatch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON:API request format"})
		return
	}
	//calls the database to execute the operation
	//if an error is detected, the code stops with its respective error
	if err := h.taskUseCase.CreateTask(&fRequestTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//generating response
	fResponseTask := tasks_response.FormatResponse(fRequestTask)
	c.JSON(http.StatusCreated, fResponseTask)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}
	fResponseTask := tasks_response.FormatResponseGetMultiple(tasks)
	c.JSON(http.StatusOK, fResponseTask)
	// Formato de respuesta en JSON:API
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := h.taskUseCase.GetTask(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error getting the values"})
		return
	}
	fResponseTask := tasks_response.FormatResponse(task)
	c.JSON(http.StatusOK, fResponseTask)

}

// función actualizar task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	fRequestTask, err := tasks_request.FormatRequestPostandPatch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fRequestTask.TaskId = int(taskID)
	if err := h.taskUseCase.UpdateTask(&fRequestTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	fResponseTask := tasks_response.FormatResponse(fRequestTask)
	c.JSON(http.StatusOK, fResponseTask)
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

}
