package modules

import (
	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/routes"
)

func init() {
	Router = gin.Default()
	mapRoutes()
}

var (
	Router *gin.Engine
)

func mapRoutes() {
  routes.TaskRoutes(Router)
}