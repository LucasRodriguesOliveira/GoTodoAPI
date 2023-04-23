package task

import (
	"log"

	"GoTodoAPI/modules/model/task"
	"GoTodoAPI/modules/gorm"
)

func Create(task *task.Schema) *task.Schema {
	result := gorm.DB.Create(task)

	if result.Error != nil {
		log.Fatalf("Failed to save Task to database")
	}

	return task
}

func Read() []task.Schema {
  var tasks []task.Schema

	result := gorm.DB.Find(&tasks)

	if result.Error != nil {
		log.Fatalf("Failed to retrieve all tasks from database")
	}

	return tasks
}

func Update(id int, oldTask *task.Schema) *task.Schema {
	oldTask.ID = uint(id)

	result := gorm.DB.Save(oldTask)

	if result.Error != nil {
		log.Fatalf("Failed to update the task (%v)", id)
	}

	return oldTask
}

func Delete(id string) {
	gorm.DB.Delete(&task.Schema{}, id)
}
