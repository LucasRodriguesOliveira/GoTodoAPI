package task

import (
	Task "GoTodoAPI/modules/model/task"
	TaskService "GoTodoAPI/modules/services/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(context *gin.Context) {
	var newTask Task.Schema

	if err := context.BindJSON(&newTask); err != nil {
		return
	}

	result := TaskService.Create(&newTask)

	context.IndentedJSON(http.StatusCreated, result)
}
