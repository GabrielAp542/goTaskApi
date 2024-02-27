package tasks_request

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestFormatRequestPostandPUT(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(sampleJSON))

	task, err := FormatRequestPostandPUT(c)

	assert.NoError(t, err)

	// Verify the fields of the returned task
	assert.Equal(t, task.Task_name, "testTask")
	assert.Equal(t, task.Completed, false)
	assert.NotEqual(t, task.Id_User, nil)

	//------fail test-----------
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)

	cf.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(badJSON))
	_, errf := FormatRequestPostandPUT(c)
	assert.Error(t, errf)
}

func TestFormatString(t *testing.T) {
	_, err := FormatString(sampleJSON)
	assert.NoError(t, err)
	_, errFail := FormatString("sampleJSON")
	assert.Error(t, errFail)
}
