package routes

import (
	"task-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/auth")
	userGroup.POST("/register", controllers.Register)  
	userGroup.POST("/login", controllers.Login)       
	userGroup.GET("/members", controllers.GetMembers) 
}
