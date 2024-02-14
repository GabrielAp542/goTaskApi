// main.go
package main

import (
	"fmt"
	"os"

	"github.com/GabrielAp542/goTask/handler/http"

	entities "github.com/GabrielAp542/goTask/internal/entities"
	repository "github.com/GabrielAp542/goTask/internal/repositories"
	usecase "github.com/GabrielAp542/goTask/internal/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Postgres conection by getting env variables
	dsn := fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s  port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	//open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//detect any error
	if err != nil {
		panic("Failed to connect to database - closing api")
	} else {
		// Migrar esquemas
		db.AutoMigrate(&entities.Task{})
		db.AutoMigrate(&entities.Users{})
		// Dependency Injection
		//inicialice the taskRepository variable by the function newTaskRepository, giving it
		//the db parameter from the conection
		taskRepository := repository.NewTaskRepository(db)
		taskUseCase := usecase.NewTaskUseCase(taskRepository)
		taskHandler := http.NewTaskHandler(*taskUseCase)

		// configuración de rutas
		router := gin.Default()
		router.GET("/tasks", taskHandler.GetTasks)
		router.GET("/tasks/:id", taskHandler.GetTask)
		router.POST("/tasks", taskHandler.CreateTask)
		router.PATCH("/tasks/:id", taskHandler.UpdateTask)
		router.DELETE("/tasks/:id", taskHandler.DeleteTask)

		// Iniciar servidor
		router.Run(":8080")
	}

}
