package http

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	// swagger embed files
	"github.com/GabrielAp542/goTaskApi/cmd/database"
	tasks_request "github.com/GabrielAp542/goTaskApi/handler/presenters/requests"
	"github.com/GabrielAp542/goTaskApi/internal/repositories"
	usecase "github.com/GabrielAp542/goTaskApi/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var requestBody = []byte(`{
    "data": {
        "attributes": {
			"task_name": "task name ex",
            "completed": true

        },
        "relationships": {
            "user": {
                "data": {
                    "id": null,
                    "type": "user"
                }
            }
        },
        "type": "tasks"
    }
}`)

// create test post
var updateBody = []byte(`{
        "data": {
            "Attributes": {
                "Task_name": "testTask_update",
                "Completed": false
            },
            "Relationships": {
                "User": {
                    "Id_User": null
                }
            }
        }
    }`)

var requestBodyWithoutName = []byte(`{
    "data": {
        "attributes": {
			"task_name": "",
            "completed": false

        },
        "relationships": {
            "user": {
                "data": {
                    "id": null,
                    "type": "user"
                }
            }
        },
        "type": "tasks"
    }
}`)

// connection variables
var db, dbF *gorm.DB
var errDB, errDBF error

func init() {
	db, errDB = database.TestingDB(false)
	dbF, errDBF = database.TestingDB(true)
}

// --testing--

func TestCreateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//---create task succesfully
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	TaskHandler.CreateTask(c)
	assert.Equal(t, http.StatusCreated, w.Code)

	//---error null body
	requestBodyf := []byte(``)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	cf.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyf))
	TaskHandler.CreateTask(cf)
	assert.Equal(t, http.StatusBadRequest, wf.Code)

	//---error no task_name
	wfNN := httptest.NewRecorder()
	cfNN, _ := gin.CreateTestContext(wfNN)
	cfNN.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyWithoutName))
	TaskHandler.CreateTask(cfNN)
	assert.Equal(t, http.StatusInternalServerError, wfNN.Code)

	//---tests bad database connections

}

func TestGetTasks(t *testing.T) {
	//testing server fail
	assert.Error(t, errDBF)
	taskRepof := repositories.NewTaskRepository(dbF)
	task_usecasef := usecase.NewTaskUseCase(taskRepof)
	TaskHandlerf := NewTaskHandler(*task_usecasef)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	TaskHandlerf.GetTasks(cf)
	assert.Equal(t, http.StatusInternalServerError, wf.Code)

	//--testing connection
	assert.NoError(t, errDB)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	TaskHandler.GetTasks(c)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetTaskId(t *testing.T) {
	//----Tests database connection
	assert.NoError(t, errDB)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//--post example
	wPT := httptest.NewRecorder()
	cPT, _ := gin.CreateTestContext(wPT)
	cPT.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	TaskHandler.CreateTask(cPT)
	assert.Equal(t, http.StatusCreated, wPT.Code)
	//formats the response from post
	format, _ := tasks_request.FormatString(wPT.Body.String())

	//--test sucessfull
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks/", nil)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: strconv.Itoa(format.TaskId)})
	TaskHandler.GetTask(c)
	log.Print(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)

	//--test bad id param
	wfidp := httptest.NewRecorder()
	cfidp, _ := gin.CreateTestContext(wfidp)
	cfidp.Request, _ = http.NewRequest("GET", "/tasks/", nil)
	cfidp.Params = append(cfidp.Params, gin.Param{Key: "id", Value: "null"})
	TaskHandler.GetTask(cfidp)
	assert.Equal(t, http.StatusBadRequest, wfidp.Code)

	//--test id not found on database
	wfidDB := httptest.NewRecorder()
	cfidDB, _ := gin.CreateTestContext(wfidDB)
	cfidDB.Request, _ = http.NewRequest("GET", "/tasks/", nil)
	cfidDB.Params = append(cfidDB.Params, gin.Param{Key: "id", Value: "100"})
	TaskHandler.GetTask(cfidDB)
	assert.Equal(t, http.StatusNotFound, wfidDB.Code)
}

// update task test
func TestUpdateTask(t *testing.T) {
	assert.NoError(t, errDB)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//--post example
	wPT := httptest.NewRecorder()
	cPT, _ := gin.CreateTestContext(wPT)
	cPT.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(updateBody))
	TaskHandler.CreateTask(cPT)
	//formats the response from post
	format, _ := tasks_request.FormatString(wPT.Body.String())

	//--test sucessfull
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/tasks/", bytes.NewBuffer(updateBody))
	c.Params = append(c.Params, gin.Param{Key: "id", Value: strconv.Itoa(format.TaskId)})
	TaskHandler.UpdateTask(c)
	assert.Equal(t, http.StatusOK, w.Code)

	//error bad id
	wfid := httptest.NewRecorder()
	cfid, _ := gin.CreateTestContext(wfid)
	cfid.Request, _ = http.NewRequest("PUT", "/tasks/", bytes.NewBuffer(requestBody))
	cfid.Params = append(cfid.Params, gin.Param{Key: "id", Value: "null"})
	TaskHandler.UpdateTask(cfid)
	assert.Equal(t, http.StatusBadRequest, wfid.Code)

	//error null body
	requestBodyf := []byte(``)
	wfN := httptest.NewRecorder()
	cfN, _ := gin.CreateTestContext(wfN)
	cfN.Request, _ = http.NewRequest("PUT", "/tasks/", bytes.NewBuffer(requestBodyf))
	cfN.Params = append(cfN.Params, gin.Param{Key: "id", Value: strconv.Itoa(format.TaskId)})
	TaskHandler.UpdateTask(cfN)
	assert.Equal(t, http.StatusBadRequest, wfN.Code)

	//test server error
	assert.Error(t, errDBF)
	taskRepof := repositories.NewTaskRepository(dbF)
	task_usecasef := usecase.NewTaskUseCase(taskRepof)
	TaskHandlerf := NewTaskHandler(*task_usecasef)
	wfs := httptest.NewRecorder()
	cfs, _ := gin.CreateTestContext(wfs)
	cfs.Request, _ = http.NewRequest("PUT", "/tasks/", bytes.NewBuffer(updateBody))
	cfs.Params = append(cfs.Params, gin.Param{Key: "id", Value: strconv.Itoa(format.TaskId)})
	TaskHandlerf.UpdateTask(cfs)
	assert.Equal(t, http.StatusInternalServerError, wfs.Code)

}

func TestDeleteTasks(t *testing.T) {
	//conection to testDB
	assert.NoError(t, errDB)
	//dependency injection testDB
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//test delete sucessfull
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	TaskHandler.DeleteTask(c)
	assert.Equal(t, http.StatusOK, w.Code)

	//test bad id error
	wfid := httptest.NewRecorder()
	cfid, _ := gin.CreateTestContext(wfid)
	cfid.Params = append(cfid.Params, gin.Param{Key: "id", Value: "null"})
	TaskHandler.DeleteTask(cfid)
	assert.Equal(t, http.StatusBadRequest, wfid.Code)

	//test server error
	assert.Error(t, errDBF)
	taskRepof := repositories.NewTaskRepository(dbF)
	task_usecasef := usecase.NewTaskUseCase(taskRepof)
	TaskHandlerf := NewTaskHandler(*task_usecasef)
	wfs := httptest.NewRecorder()
	cfs, _ := gin.CreateTestContext(wfs)
	cfs.Params = append(cfs.Params, gin.Param{Key: "id", Value: "1"})
	TaskHandlerf.DeleteTask(cfs)
	assert.Equal(t, http.StatusInternalServerError, wfs.Code)

}
