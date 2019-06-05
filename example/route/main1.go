package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

type Hander struct {
	http.Handler
}
type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerInterface interface {
	Init(ct *Context, cn string) // 初始化上下文和子类名称
	Prepare()                    // 开始执行之前的一些处理
	Get()                        // method=GET 的处理
	Post()                       // method=POST 的处理
	Delete()                     // method=DELETE 的处理
	Put()                        // method=PUT 的处理
	Head()                       // method=HEAD 的处理
	Patch()                      // method=PATCH 的处理
	Options()                    // method=OPTIONS 的处理
	Finish()                     // 执行完成之后的处理
	Render() error               // 执行完 method 对应的方法之后渲染页面
}
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         map[string]string
}

func (ctx *Context) Redirect(status int, localurl string) {
	http.Redirect(ctx.ResponseWriter, ctx.Request, localurl, status)
}

type App struct {
	http.Handler
	routers      []*controllerInfo
	ViewPath     string
	RecoverPanic bool
	AutoRender   bool
}

func (a *App) AddRoute(pattern string, c ControllerInterface) {
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
	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	a.routers = append(a.routers, route)

}

type Controller struct {
	App       *App
	Ct        *Context
	Tpl       *template.Template
	Data      map[interface{}]interface{}
	ChildName string
	TplNames  string
	Layout    []string
	TplExt    string
}

func (c *Controller) Init(ct *Context, cn string) {
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.ChildName = cn
	c.Ct = ct
	c.TplExt = "tpl"
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Render() error {
	if len(c.Layout) > 0 {
		var filenames []string
		for _, file := range c.Layout {
			filenames = append(filenames, path.Join(c.App.ViewPath, file))
		}
		t, err := template.ParseFiles(filenames...)
		if err != nil {
			log.Fatal("template ParseFiles err:", err)
		}
		err = t.ExecuteTemplate(c.Ct.ResponseWriter, c.TplNames, c.Data)
		if err != nil {
			log.Fatal("template Execute err:", err)
		}
	} else {
		if c.TplNames == "" {
			c.TplNames = c.ChildName + "/" + c.Ct.Request.Method + "." + c.TplExt
		}
		t, err := template.ParseFiles(path.Join(c.App.ViewPath, c.TplNames))
		if err != nil {
			log.Fatal("template ParseFiles err:", err)
		}
		err = t.Execute(c.Ct.ResponseWriter, c.Data)
		if err != nil {
			log.Fatal("template Execute err:", err)
		}
	}
	return nil
}

type UserController struct {
	Controller
}

func (c *UserController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "\nHello World!")
}
func (c *Controller) Redirect(url string, code int) {
	c.Ct.Redirect(code, url)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			if !app.RecoverPanic {
				// go back to panic
				panic(err)
			} else {
				log.Fatal("Handler crashed with error", err)
				for i := 1; ; i += 1 {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					log.Fatal(file, line)
				}
			}
		}
	}()
	started := false
	requestPath := r.URL.Path

	//find a matching Route
	for _, route := range app.routes {

		// check if Route pattern matches url
		if !route.regex.MatchString(requestPath) {
			continue
		}

		// get submatches (params)
		matches := route.regex.FindStringSubmatch(requestPath)

		// double check that the Route matches the URL pattern.
		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := make(map[string]string)
		if len(route.params) > 0 {
			// add url parameters to the query param map
			values := r.URL.Query()
			for i, match := range matches[1:] {
				values.Add(route.params[i], match)
				params[route.params[i]] = match
			}

			// reassemble query params and add to RawQuery
			r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
			// r.URL.RawQuery = url.Values(values).Encode()
		}

		// Invoke the request handler
		vc := reflect.New(route.controllerType)

		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		ct := &Context{ResponseWriter: w, Request: r, Params: params}
		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())

		init.Call(in)
		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)
		if r.Method == "GET" {
			method = vc.MethodByName("Get")
			method.Call(in)
		} else if r.Method == "POST" {
			method = vc.MethodByName("Post")
			method.Call(in)
		} else if r.Method == "HEAD" {
			method = vc.MethodByName("Head")
			method.Call(in)
		} else if r.Method == "DELETE" {
			method = vc.MethodByName("Delete")
			method.Call(in)
		} else if r.Method == "PUT" {
			method = vc.MethodByName("Put")
			method.Call(in)
		} else if r.Method == "PATCH" {
			method = vc.MethodByName("Patch")
			method.Call(in)
		} else if r.Method == "OPTIONS" {
			method = vc.MethodByName("Options")
			method.Call(in)
		}
		if app.AutoRender {
			method = vc.MethodByName("Render")
			method.Call(in)
		}
		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	// if no matches to url, throw a not found exception
	if started == false {
		http.NotFound(w, r)
	}

	// fmt.Fprintf(w, r.URL.Path)
	// fmt.Fprintf(w, "\nHello World!") // 这个写入到 w 的是输出到客户端的
}

func main() {

	app := &App{
		ViewPath:     "./",
		RecoverPanic: false,
		AutoRender:   false,
	}

	app.AddRoute("/users", &UserController{})
	app.AddRoute("/users/:id", &UserController{})
	err := http.ListenAndServe(":9090", app) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
