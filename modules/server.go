package modules

import (
	GO_ENV "GoTodoAPI/modules/config"
	GORM   "GoTodoAPI/modules/gorm"

	r "GoTodoAPI/modules/routes"
)

func Server() {
	GO_ENV.Load()
	GORM.StartUp()

	r.Router.Run(GO_ENV.APP["host"])
}
