package tasks_response

import (
	"testing"

	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestFormatResponse(t *testing.T) {

	taskOk := entities.Task{
		TaskId:    1,
		Task_name: "testTask",
		Completed: true,
		Id_User:   nil,
	}

	formatedOk := FormatResponse(taskOk)
	if formatedOk == nil {
		t.Error("formated not returned")
	}
}

func TestFormatResponseGetMultiple(t *testing.T) {

	taskOk := []entities.Task{
		{
			TaskId:    1,
			Task_name: "testTask1",
			Completed: true,
			Id_User:   nil,
		},
		{
			TaskId:    2,
			Task_name: "testTask2",
			Completed: true,
			Id_User:   nil,
		},
		{
			TaskId:    3,
			Task_name: "testTask3",
			Completed: false,
			Id_User:   nil,
		},
	}

	formatedOk := FormatResponseGetMultiple(taskOk)
	assert.NotNil(t, formatedOk)
}
