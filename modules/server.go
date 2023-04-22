package modules

import (
	U      "GoTodoAPI/util"
	GO_ENV "GoTodoAPI/modules/config"
	GORM   "GoTodoAPI/modules/gorm"
)

func Server() {
	GO_ENV.LoadEnv()
	GORM.StartUp()
	host := GO_ENV.ServerHost()
	port, err := GO_ENV.ServerPort()

	if err != nil {
		return
	}

	path := U.Addr(host, port)

	Router.Run(path)
}
