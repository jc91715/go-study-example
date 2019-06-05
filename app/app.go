package app

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	http.Handler
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "\nHello World!") // 这个写入到 w 的是输出到客户端的
}
func Run() {
	app := &App{}
	err := http.ListenAndServe(":9090", app) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
