package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	TaskService "GoTodoAPI/modules/services/task"
)

func list(c *gin.Context) {
	tasks := TaskService.Read()

	c.IndentedJSON(http.StatusOK, tasks)
}
