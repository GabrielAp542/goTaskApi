package presenters

import (
	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/gin-gonic/gin"
)

// estruct to decode post
type TaskData struct {
	Data struct {
		Attributes struct {
			Task_name string `json:"task_name"`
			Completed bool   `json:"completed"`
		} `json:"attributes"`
		Relationships struct {
			User struct {
				Id_User *int `json:"id"`
			} `json:"user"`
		} `json:"relationships"`
	} `json:"data"`
}

func FormatResponse(c *gin.Context, task *entities.Task, status int) {
	//status := http.StatusCreated
	c.JSON(status, gin.H{
		"data": gin.H{
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
		},
	})
}
