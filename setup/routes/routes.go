package routes

import (
	"github.com/GabrielAp542/goTaskApi/handler/http"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(handler *http.TaskHandler) *gin.Engine {
	router := gin.Default()

	router.Use(handler.PolicyEnferocer)
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasks/:id", handler.GetTask)
	router.POST("/tasks", handler.CreateTask)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)
	return router

}
