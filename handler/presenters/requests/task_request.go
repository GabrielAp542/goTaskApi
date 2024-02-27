package tasks_request

import (
	"encoding/json"

	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/gin-gonic/gin"
)

// estruct to decode post
type TaskDataDoc struct {
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

type ErrorDoc struct {
	Error string `json:"error"`
}

func FormatRequestPostandPUT(c *gin.Context) (entities.Task, error) {

	newTask := &entities.Task{}
	err := c.ShouldBindJSON(&TaskData)
	if err != nil {
		return *newTask, err
	}
	newTask.TaskId = TaskData.Data.TaskId
	newTask.Task_name = TaskData.Data.Attributes.Task_name
	newTask.Completed = TaskData.Data.Attributes.Completed
	newTask.Id_User = TaskData.Data.Relationships.User.Id_User
	return *newTask, err

}

func FormatString(format string) (entities.Task, error) {

	formatedTask := &entities.Task{}
	err := json.Unmarshal([]byte(format), &TaskData)
	if err != nil {
		return *formatedTask, err
	}
	formatedTask.TaskId = TaskData.Data.TaskId
	formatedTask.Task_name = TaskData.Data.Attributes.Task_name
	formatedTask.Completed = TaskData.Data.Attributes.Completed
	formatedTask.Id_User = TaskData.Data.Relationships.User.Id_User
	return *formatedTask, err

}
