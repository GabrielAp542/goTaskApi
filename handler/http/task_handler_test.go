package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielAp542/goTask/internal/entities"
	"github.com/GabrielAp542/goTask/internal/repositories"
	usecase "github.com/GabrielAp542/goTask/internal/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "host=172.18.0.2 user=postgres password=1234 dbname=test_tasksDB port=5432"
	// Postgres conection by getting env variables
	//open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//detect any error
	if err != nil {
		panic("Failed to connect to database - closing api")
	}
	// Migrar esquemas
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Users{})
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTestDB()
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	router := gin.Default()
	router.POST("/tasks", TaskHandler.CreateTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
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
	c.Request = req
	TaskHandler.CreateTask(c)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	//error null body
	requestBodyf := []byte(``)
	wf := httptest.NewRecorder()
	cf, _ := gin.CreateTestContext(wf)
	reqf, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyf))
	cf.Request = reqf
	TaskHandler.CreateTask(cf)
	if wf.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, wf.Code)
	}

	requestBodyNoName := []byte(`{
		        "data": {
		            "Attributes": {
		                "Task_name": "",
		                "Completed": false
		            },
		            "Relationships": {
		                "User": {
		                    "Id_User": null
		                }
		            }
		        }
		    }`)
	//error no name
	wfNN := httptest.NewRecorder()
	cfNN, _ := gin.CreateTestContext(wfNN)
	reqNoName, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBodyNoName))
	cfNN.Request = reqNoName
	TaskHandler.CreateTask(cfNN)
	if wfNN.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, wfNN.Code)
	}

}

func TestGetTasks(t *testing.T) {
	db := setupTestDB()
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	TaskHandler.GetTasks(c)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetTaskId(t *testing.T) {
	//dependency injection
	db := setupTestDB()
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
	db := setupTestDB()
	taskRepo := repositories.NewTaskRepository(db)
	task_usecase := usecase.NewTaskUseCase(taskRepo)
	TaskHandler := NewTaskHandler(*task_usecase)

	router := gin.Default()
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

	router.PATCH("/tasks/:id", TaskHandler.UpdateTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("PATCH", "/tasks/1", bytes.NewBuffer(updateBody))
	c.Request = req
	TaskHandler.UpdateTask(c)
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

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
	db := setupTestDB()
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
	if wfid.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, wfid.Code)
	}
}
