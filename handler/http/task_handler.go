// delivery/http/task_handler.go
package http

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	mappermissions "github.com/GabrielAp542/goTaskApi/handler/map_permissions"
	tasks_request "github.com/GabrielAp542/goTaskApi/handler/presenters/requests"
	tasks_response "github.com/GabrielAp542/goTaskApi/handler/presenters/responses"
	usecase "github.com/GabrielAp542/goTaskApi/internal/usecases"
	"github.com/Nerzal/gocloak/v13"

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

// middleware auth
func (h *TaskHandler) PolicyEnferocer(c *gin.Context) {

	//path := c.FullPath()
	//method := c.Request.Method

	scope := mappermissions.ConfigPermissions[c.FullPath()][c.Request.Method]
	if scope == "public" {
		c.Next()
		return
	}
	//scope := path[c.Request.Method]

	kc_client := gocloak.NewClient(os.Getenv("KC_URL"))
	ctx := context.Background()

	//gets token from headers
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "missing token"})
		c.Abort()
		return
	}

	accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	//evaluate token
	/*result, err := kc_client.RetrospectToken(ctx, accessToken, os.Getenv("KC_CLIENT_ID"), os.Getenv("KC_CLIENT_SECRET"), os.Getenv("KC_REALM"))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not valid token"})
		c.Abort()
		return
	}
	if !*result.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": "not active token"})
		c.Abort()
		return
	}
	*/
	clientID := os.Getenv("KC_CLIENT_ID")

	options := gocloak.RequestingPartyTokenOptions{
		Audience:    &clientID,
		Permissions: &[]string{"todo_tasks#" + scope},
	}
	permissions, err := kc_client.GetRequestingPartyPermissionDecision(ctx, accessToken, os.Getenv("KC_REALM"), options)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied or auth server not found"})
		c.Abort()
		return
	}
	if permissions.Result == nil || !*permissions.Result {
		c.JSON(http.StatusForbidden, gin.H{"status": "access denied"})
		c.Abort()
		return
	}
}
