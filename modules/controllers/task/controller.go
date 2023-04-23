package task

import "GoTodoAPI/shared"

var Controller shared.Controller

func init() {
	Controller = shared.Controller{
		BasePath: "task",
		Actions: []shared.Action{
			{ Path: "", Method: shared.POST, Function: create },
			{ Path: "", Method: shared.GET, Function: list },
			{ Path: ":id", Method: shared.PUT, Function: update },
			{ Path: ":id", Method: shared.DELETE, Function: delete },
		},
	}
}
