package routes

import (
	"task-management/controllers"
	"task-management/middlewares"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	taskGroup.Use(middlewares.AuthMiddleware()) 

	taskGroup.POST("/", controllers.CreateTask)  
	taskGroup.GET("/", controllers.GetTasks)     
	taskGroup.PUT("/:id", controllers.UpdateTask) 
	taskGroup.DELETE("/:id", controllers.DeleteTask)
}
