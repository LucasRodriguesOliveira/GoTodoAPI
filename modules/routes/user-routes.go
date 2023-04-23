package routes

import (
	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/controllers/user"
)

func userRoutes(router *gin.Engine) {
	userRoutes := router.Group(user.Controller.BasePath)

	addRoutes(userRoutes, user.Controller.Actions)
}
