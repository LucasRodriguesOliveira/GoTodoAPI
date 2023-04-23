package gorm

import (
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	GO_ENV "GoTodoAPI/modules/config"

	"GoTodoAPI/modules/model/task"
	"GoTodoAPI/modules/model/user"
)

var (
	DB *gorm.DB
)

func StartUp() {
  var dsn string = GO_ENV.DSN

	conn, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	conn.AutoMigrate(
		&task.Schema{},
		&user.Schema{},
	)

	DB = conn
}
