// delivery/http/task_handler.go
package http

import (
	"net/http"
	"strconv"

	tasks_request "github.com/GabrielAp542/goTask-Api-Gabriel/handler/presenters/requests"
	tasks_response "github.com/GabrielAp542/goTask-Api-Gabriel/handler/presenters/responses"
	usecase "github.com/GabrielAp542/goTask-Api-Gabriel/internal/usecases"

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

// @Router /tasks [post]
// @Summary Create a task
// @Description Creates a task to database
// @Tags tasks
// @accept json
// @Produce json
// @Param Tasks body tasks_request.TaskDataDoc true "Task body"
// @Success 201 {object} tasks_request.TaskDataDoc "Task created"
// @Failure 400 {object} tasks_request.ErrorDoc "Invalid JSON:API request format"
// @Failure 500 {object} tasks_request.ErrorDoc "Internal Server Error"
func (h *TaskHandler) CreateTask(c *gin.Context) {
	fRequestTask, err := tasks_request.FormatRequestPostandPUT(c)
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

// @Router /tasks [get]
// @Summary get all tasks registrated
// @Description gets all the tasks from database
// @Tags tasks
// @Produce json
// @Success 200 {object} tasks_request.TaskDataDoc "Task created"
// @Failure 400 {object} tasks_request.ErrorDoc "Internal Server Error"
// @Failure 500 {object} tasks_request.ErrorDoc "Internal Server Error"
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

// @Router /tasksId/{id} [get]
// @Summary Get task id
// @Description gets an specific task from database
// @Tags tasks
// @Produce json
// @Param id path int true "task id"
// @Success 200
// @Failure 400 {object} tasks_request.TaskDataDoc "Invalid task ID"
// @Failure 404 {object} tasks_request.TaskDataDoc "value not fould"
func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := h.taskUseCase.GetTask(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Value not fould"})
		return
	}
	fResponseTask := tasks_response.FormatResponse(task)
	c.JSON(http.StatusOK, fResponseTask)

}

// función actualizar task

// @Router /tasks/{id} [put]
// @Summary update task
// @Description update a specific task from database
// @Tags tasks
// @Produce json
// @Param id path int true "task id"
// @Param Tasks body tasks_request.TaskDataDoc true "Task body"
// @Success 200 {object} tasks_request.TaskDataDoc "updated task"
// @Failure 400 {object} tasks_request.TaskDataDoc "Invalid task ID or Body"
// @Failure 404 {object} tasks_request.TaskDataDoc "server Error"
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	fRequestTask, err := tasks_request.FormatRequestPostandPUT(c)
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

// función eliminar task

// @Router /tasks/{id} [delete]
// @Summary delete task
// @Description delete an specific task from database
// @Tags tasks
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} tasks_request.TaskDataDoc "deleted task"
// @Failure 400 {object} tasks_request.TaskDataDoc "Invalid task ID "
// @Failure 404  "server Error"
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
