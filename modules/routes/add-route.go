package routes

import (
	"GoTodoAPI/shared"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup, actions []shared.Action) {
	for _, action := range actions {
		addRoute(rg, action)
	}
}

func addRoute(rg *gin.RouterGroup, action shared.Action) {
	switch action.Method {
	case shared.GET:
		rg.GET(action.Path, action.Function)
	case shared.POST:
		rg.POST(action.Path, action.Function)
	case shared.PUT:
		rg.PUT(action.Path, action.Function)
	case shared.DELETE:
		rg.DELETE(action.Path, action.Function)
	}
}
