package presenters

import (
	"log"

	entities "github.com/GabrielAp542/goTask/internal/entities"
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

func DecodeTask(c *gin.Context) {
	var taskData TaskData
	var task entities.Task
	c.BindJSON(&taskData)
	log.Print(taskData.Data.Attributes.Task_name)
	task.Task_name = taskData.Data.Attributes.Task_name
	log.Print(&task.Task_name)
}
