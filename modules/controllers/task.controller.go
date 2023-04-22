package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/model"
	TaskService "GoTodoAPI/modules/services"
)

func CreateTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	result := TaskService.CreateTask(&newTask)

	c.IndentedJSON(http.StatusCreated, result)
}

func FindTaskById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}
