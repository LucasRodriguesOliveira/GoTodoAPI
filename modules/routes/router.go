package routes

import (
	"github.com/gin-gonic/gin"
)

func init() {
	Router = gin.Default()
	mapRoutes()
}

var (
	Router *gin.Engine
)

func mapRoutes() {
  taskRoutes(Router)
	userRoutes(Router)
}