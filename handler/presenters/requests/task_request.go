package tasks_request

import (
	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/gin-gonic/gin"
)

// estruct to decode post
var TaskData struct {
	Data struct {
		TaskId     int `gorm:"primaryKey" json:"task_id"`
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

func FormatRequestPostandPatch(c *gin.Context) (entities.Task, error) {

	newTask := &entities.Task{}
	err := c.ShouldBindJSON(&TaskData)
	newTask.TaskId = TaskData.Data.TaskId
	newTask.Task_name = TaskData.Data.Attributes.Task_name
	newTask.Completed = TaskData.Data.Attributes.Completed
	newTask.Id_User = TaskData.Data.Relationships.User.Id_User
	return *newTask, err
}
