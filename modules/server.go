package modules

import (
	GO_ENV "GoTodoAPI/modules/config"
	GORM   "GoTodoAPI/modules/gorm"
)

func Server() {
	GO_ENV.Load()
	GORM.StartUp()

	Router.Run(GO_ENV.APP["host"])
}
