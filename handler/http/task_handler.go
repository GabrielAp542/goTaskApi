// delivery/http/task_handler.go
package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GabrielAp542/goTask/handler/formats"
	"github.com/GabrielAp542/goTask/handler/presenters"
	"github.com/GabrielAp542/goTask/internal/entities"
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

	if err := c.ShouldBindJSON(&formats.JsonFormat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON:API request format"})
		return
	}

	// Crear una nueva tarea desde los datos proporcionados
	newTask := &entities.Task{
		Task_name: formats.JsonFormat.Data.Attributes.Task_name,
		Completed: formats.JsonFormat.Data.Attributes.Completed,
	}

	//calls the database to execute the operation
	//if an error is detected, the code stops with its respective error
	if err := h.taskUseCase.CreateTask(newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//generating response
	presenters.FormatResponse(c, newTask, http.StatusCreated)
	//confirmation of the success of the operation
}
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}
	// Formato de respuesta en JSON:API
	var responseData []gin.H
	for _, task := range tasks {
		responseData = append(responseData,
			gin.H{
				"type": "tasks",
				"id":   task.TaskId,
				"attributes": gin.H{
					"task_name": task.Task_name,
					"completed": task.Completed,
				},
				"relationships": gin.H{
					"user": gin.H{
						"data": gin.H{
							"type": "user",
							"id":   task.Id_User,
						},
					},
				},
				"links": gin.H{
					"self": fmt.Sprintf("http://localhost:8080/tasks/%s", strconv.Itoa(task.TaskId)),
				},
			})
	}

	c.JSON(http.StatusOK,
		gin.H{
			"links": gin.H{},
			"data":  responseData,
		})
	//c.JSON(http.StatusOK, tasks)
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
	c.JSON(http.StatusOK,
		gin.H{
			"links": gin.H{
				"self": fmt.Sprintf("http://localhost:8080/tasks/%s", strconv.Itoa(task.TaskId)),
			},
			"data": gin.H{
				"type": "tasks",
				"id":   task.TaskId,
				"attributes": gin.H{
					"task_name": task.Task_name,
					"completed": task.Completed,
				},
				"relationships": gin.H{
					"user": gin.H{
						"links": gin.H{
							"self": fmt.Sprintf("http://localhost:8080/users/%s", strconv.Itoa(task.User.UserId)),
						},
						"data": gin.H{
							"type": "user",
							"id":   task.Id_User,
						},
					},
				},
			},
		})
}

// función actualizar task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := c.ShouldBindJSON(&formats.JsonFormat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskUpdate := entities.Task{
		Task_name: formats.JsonFormat.Data.Attributes.Task_name,
		Completed: formats.JsonFormat.Data.Attributes.Completed,
		Id_User:   formats.JsonFormat.Data.Relationships.User.Id_User,
	}

	taskUpdate.TaskId = int(taskID)
	if err := h.taskUseCase.UpdateTask(&taskUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"links": gin.H{
			"self": fmt.Sprintf("http://localhost:8080/tasks/%s", strconv.Itoa(taskUpdate.TaskId)),
		},
		"data": gin.H{
			"type": "tasks",
			"id":   taskUpdate.TaskId,
			"attributes": gin.H{
				"task_name": taskUpdate.Task_name,
				"completed": taskUpdate.Completed,
			},
			"relationships": gin.H{
				"user": gin.H{
					"links": gin.H{
						"self": fmt.Sprintf("http://localhost:8080/users/%s", strconv.Itoa(taskUpdate.User.UserId)),
					},
					"data": gin.H{
						"type": "user",
						"id":   taskUpdate.Id_User,
					},
				},
			},
		},
	})
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

	c.JSON(http.StatusNoContent, gin.H{"result": "Invalid task ID"})
}
