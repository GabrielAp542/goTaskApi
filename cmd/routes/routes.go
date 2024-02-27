package routes

import (
	"github.com/GabrielAp542/goTask/handler/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRoutes(handler *http.TaskHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasksId/:id", handler.GetTask)
	router.POST("/tasks", handler.CreateTask)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)
	return router

}
