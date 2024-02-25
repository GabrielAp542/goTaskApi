package tasks_response

import (
	"fmt"
	"strconv"

	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/gin-gonic/gin"
)

// estruct to decode post

func FormatResponse(task entities.Task) gin.H {
	//status := http.StatusCreated
	formated := gin.H{
		"data": gin.H{
			"type":    "tasks",
			"task_id": task.TaskId,
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
		},
	}
	return formated
}

func FormatResponseGetMultiple(tasks []entities.Task) gin.H {
	var responseData []gin.H
	//status := http.StatusCreated
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

	formated := gin.H{
		"links": gin.H{},
		"data":  responseData,
	}
	return formated
}
