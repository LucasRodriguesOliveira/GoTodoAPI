package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"GoTodoAPI/modules/model"
	TaskService "GoTodoAPI/modules/services"
)

func CreateTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	result := TaskService.Create(&newTask)

	c.IndentedJSON(http.StatusCreated, result)
}

func ListTasks(c *gin.Context) {
	tasks := TaskService.Read()

	c.IndentedJSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
  var task model.Task

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	updatedTask := TaskService.Update(id, &task)

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	TaskService.Delete(id)
}
