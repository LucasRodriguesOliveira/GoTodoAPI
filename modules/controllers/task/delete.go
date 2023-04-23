package task

import (
	"github.com/gin-gonic/gin"

	TaskService "GoTodoAPI/modules/services/task"
)

func delete(c *gin.Context) {
	id := c.Param("id")
	TaskService.Delete(id)
}
