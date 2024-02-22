package routes

import (
	handlers "github.com/GabrielAp542/goTask/handler/http"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(handler *handlers.TaskHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasks/:id", handler.GetTask)
	router.POST("/tasks", handler.CreateTask)
	router.PATCH("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)
	return router

}
