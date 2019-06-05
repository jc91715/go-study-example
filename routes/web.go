package routes

import (
	"controller"
	"reflect"
	"regexp"
)

type Route struct {
	Regex          *regexp.Regexp
	Methods        map[string]string
	Params         map[int]string
	ControllerType reflect.Type
}

type app interface {
	AddRoute(pattern string, m map[string]string, c controller.ControllerInterface)
}

func AddRoutes(a app) {
	a.AddRoute("/users", map[string]string{
		"GET": "Index",
	}, &controller.UserController{})

	a.AddRoute("/users/:id", map[string]string{
		"GET": "Show",
	}, &controller.UserController{})
}
