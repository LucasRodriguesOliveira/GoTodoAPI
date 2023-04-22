package gorm

import (
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	GO_ENV "GoTodoAPI/modules/config"
	"GoTodoAPI/modules/model"
)

var (
	DB *gorm.DB
)

func StartUp() {
  var dsn string = GO_ENV.LoadDBEnv()["dsn"]

	conn, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	conn.AutoMigrate(&model.Task{})

	DB = conn
}
