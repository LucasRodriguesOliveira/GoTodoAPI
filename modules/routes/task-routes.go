package routes

import (
	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/controllers/task"
)

func taskRoutes(router *gin.Engine) {
	taskRoutes := router.Group(task.Controller.BasePath)

	addRoutes(taskRoutes, task.Controller.Actions)
}
