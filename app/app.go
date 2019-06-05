package app

import (
	"controller"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"routes"
	"strings"
)

type App struct {
	http.Handler
	routes []*routes.Route
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "\nHello World!") // 这个写入到 w 的是输出到客户端的
}

func (a *App) AddRoutes() {
	routes.AddRoutes(a)
}
func (a *App) AddRoute(pattern string, c controller.ControllerInterface) {
	parts := strings.Split(pattern, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			// a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	// recreate the url pattern, with parameters replaced
	// by regular expressions. then compile the regex

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		// TODO add error handling here to avoid panic
		panic(regexErr)
		return
	}

	// now create the Route
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &routes.Route{}
	route.Regex = regex
	route.Params = params
	route.ControllerType = t

	a.routes = append(a.routes, route)

}

func Run() {
	app := &App{}
	// fmt.Println(routes)
	app.AddRoutes()
	err := http.ListenAndServe(":9090", app) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
