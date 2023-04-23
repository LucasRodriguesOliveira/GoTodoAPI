package task

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	Task "GoTodoAPI/modules/model/task"
	TaskService "GoTodoAPI/modules/services/task"
)

func update(context *gin.Context) {
	var task Task.Schema

	if err := context.BindJSON(&task); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		log.Fatalf("Error while trying to parser the id")
	}

	updatedTask := TaskService.Update(id, &task)

	context.IndentedJSON(http.StatusOK, updatedTask)
}
