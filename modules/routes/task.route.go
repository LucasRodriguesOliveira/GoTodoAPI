package routes

import (
	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/controllers"
)

func TaskRoutes(router *gin.Engine) {
	router.POST("/task", controllers.CreateTask)
	router.GET("/task", controllers.ListTasks)
	router.PUT("/task/:id", controllers.UpdateTask)
	router.DELETE("/task/:id", controllers.DeleteTask)
}
