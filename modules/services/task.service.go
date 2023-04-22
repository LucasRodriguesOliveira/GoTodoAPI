package services

import (
	"log"

	"GoTodoAPI/modules/model"
	"GoTodoAPI/modules/gorm"
)

func Create(task *model.Task) *model.Task {
	result := gorm.DB.Create(task)

	if result.Error != nil {
		log.Fatalf("Failed to save Task to database")
	}

	return task
}

func Read() []model.Task {
  var tasks []model.Task

	result := gorm.DB.Find(&tasks)

	if result.Error != nil {
		log.Fatalf("Failed to retrieve all tasks from database")
	}

	return tasks
}

func Update(id int, oldTask *model.Task) *model.Task {
	oldTask.ID = uint(id)

	result := gorm.DB.Save(oldTask)

	if result.Error != nil {
		log.Fatalf("Failed to update the task (%v)", id)
	}

	return oldTask
}

func Delete(id string) {
	gorm.DB.Delete(&model.Task{}, id)
}
