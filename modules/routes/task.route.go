package routes

import (
	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/controllers"
)

func TaskRoutes(router *gin.Engine) {
	router.POST("/task", controllers.CreateTask)
}
