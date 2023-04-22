package services

import (
	"log"

	"GoTodoAPI/modules/model"
	"GoTodoAPI/modules/gorm"
)

func CreateTask(task *model.Task) *model.Task {
	result := gorm.DB.Create(task)

	if result.Error != nil {
		log.Fatalf("Failed to save Task to database")
	}

	return task
}
