package user

import "GoTodoAPI/shared"

var Controller shared.Controller

func init() {
	Controller = shared.Controller{
		BasePath: "user",
		Actions: []shared.Action{
			{},
		},
	}
}