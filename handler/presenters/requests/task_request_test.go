package tasks_request

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

const sampleJSON = `{
    "data": {
        "Attributes": {
            "TaskId": "123",
            "Task_name": "testTask",
            "Completed": false
        },
        "Relationships": {
            "User": {
                "Id_User": null
            }
        }
    }
}`

const badJSON = `{
    name: "badTask",
	compleated: false
}`

func TestFormatRequestPostandPatch(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(sampleJSON))

	task, err := FormatRequestPostandPatch(c)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the fields of the returned task
	if task.Task_name != "testTask" || task.Completed != false || task.Id_User != nil {
		t.Errorf("task format not correct: %+v", task)
	}

	//------fail test-----------
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)

	cf.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(badJSON))
	_, errf := FormatRequestPostandPatch(c)
	if errf == nil {
		t.Errorf("error not recived but expected: %v", errf)
	}
}
