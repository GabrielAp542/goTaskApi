package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielAp542/goTask/cmd/database"
	"github.com/GabrielAp542/goTask/internal/repositories"
	usecase "github.com/GabrielAp542/goTask/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var requestBody = []byte(`{
	"data": {
		"Attributes": {
			"Task_name": "testTask",
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
		"Attributes": {
			"Task_name": "testTask",
			"Completed": false
		},
		"Relationships": {
			"User": {
				"Id_User": null
			}
		}
	}
}`)

func setupTestDB() (*gorm.DB, error) {
	g, err := database.Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	return g, err
}

func setupTestDBFail() (*gorm.DB, error) {
	g, err := database.Conection("uwu",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	return g, err
}

func TestCreateTask(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	router := gin.Default()
	router.POST("/tasks", TaskHandler.CreateTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	c.Request = req
	TaskHandler.CreateTask(c)
	assert.Equal(t, http.StatusCreated, w.Code)

	//error null body
	requestBodyf := []byte(``)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	reqf, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyf))
	cf.Request = reqf
	TaskHandler.CreateTask(cf)
	assert.Equal(t, http.StatusBadRequest, wf.Code)
	//error no name
	wfNN := httptest.NewRecorder()
	cfNN, _ := gin.CreateTestContext(wfNN)
	reqNoName, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyWithoutName))
	cfNN.Request = reqNoName
	TaskHandler.CreateTask(cfNN)
	assert.Equal(t, http.StatusInternalServerError, wfNN.Code)

}

func TestGetTasks(t *testing.T) {

	//testing fail connection
	dbf, errf := setupTestDBFail()
	assert.Error(t, errf)
	taskRepof := repositories.NewTaskRepository(dbf)
	task_usecasef := usecase.NewTaskUseCase(taskRepof)
	TaskHandlerf := NewTaskHandler(*task_usecasef)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	TaskHandlerf.GetTasks(cf)

	assert.Equal(t, http.StatusInternalServerError, wf.Code)

	db, err := setupTestDB()
	assert.NoError(t, err)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	TaskHandler.GetTasks(c)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetTaskId(t *testing.T) {
	//dependency injection
	db, errDB := database.Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.NoError(t, errDB)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//post example
	router := gin.Default()
	router.POST("/tasks", TaskHandler.CreateTask)
	wP := httptest.NewRecorder()
	cP, _ := gin.CreateTestContext(wP)
	requestBody := []byte(`{
	        "data": {
	            "Attributes": {
	                "Task_name": "testTask",
	                "Completed": false
	            },
	            "Relationships": {
	                "User": {
	                    "Id_User": null
	                }
	            }
	        }
	    }`)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
	cP.Request = req
	TaskHandler.CreateTask(cP)
	if wP.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, wP.Code)
	}
	//test get 1
	//router.GET("/tasks/:id", TaskHandler.GetTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	//req, err := http.NewRequest("GET", "/tasks/1", nil)
	/*log.Print(err)
	c.Request = req*/

	TaskHandler.GetTask(c)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	//test get error id
	wfid := httptest.NewRecorder()
	cfid, _ := gin.CreateTestContext(wfid)
	cfid.Params = append(cfid.Params, gin.Param{Key: "id", Value: "null"})
	TaskHandler.GetTask(cfid)
	if wfid.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, wfid.Code)
	}
}

// update task test
func TestUpdateTask(t *testing.T) {
	db, err := database.Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.NoError(t, err)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	//router := gin.Default()
	//create test post
	updateBody := []byte(`{
        "data": {
            "Attributes": {
                "Task_name": "testTask",
                "Completed": false
            },
            "Relationships": {
                "User": {
                    "Id_User": null
                }
            }
        }
    }`)

	//router.PATCH("/tasks/:id", TaskHandler.UpdateTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("PATCH", "/tasks/", bytes.NewBuffer(updateBody))
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	c.Request = req
	TaskHandler.UpdateTask(c)
	assert.Equal(t, http.StatusOK, w.Code)

	//error null body
	requestBodyf := []byte(``)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	reqf, _ := http.NewRequest("PATCH", "/tasks/1", bytes.NewBuffer(requestBodyf))
	cf.Request = reqf
	TaskHandler.CreateTask(cf)
	if wf.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, wf.Code)
	}
}

func TestDeleteTasks(t *testing.T) {
	db, err := database.Conection("172.18.0.2",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.NoError(t, err)
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	TaskHandler.DeleteTask(c)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	//test get error id
	wfid := httptest.NewRecorder()
	cfid, _ := gin.CreateTestContext(wfid)
	cfid.Params = append(cfid.Params, gin.Param{Key: "id", Value: "null"})
	TaskHandler.DeleteTask(cfid)
	assert.Equal(t, http.StatusBadRequest, wfid.Code)

	//test server error
	_, errDB := database.Conection("asad",
		"postgres",
		"1234",
		"test_tasksDB",
		"5432")
	assert.NoError(t, errDB)
	wfs := httptest.NewRecorder()
	cfs, _ := gin.CreateTestContext(wfs)
	cfs.Params = append(cfs.Params, gin.Param{Key: "id", Value: "1"})
	TaskHandler.DeleteTask(cfs)
	assert.Equal(t, http.StatusInternalServerError, wfs.Code)

}
