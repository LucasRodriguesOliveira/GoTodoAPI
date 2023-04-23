package shared

import "github.com/gin-gonic/gin"

type Controller struct {
	BasePath string
	Actions  []Action
}

type Action struct {
	Path     string
	Method   int
	Function func(*gin.Context)
}

const (
	GET    int = 0
	POST   int = 1
	PUT    int = 2
	DELETE int = 3
	PATCH  int = 4
)
