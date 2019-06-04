package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hander struct {
	http.Handler
}

func (h *Hander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)       //可根据Path 分配请求到相对应的控制器
	fmt.Fprintf(w, "\nHello World!") // 这个写入到 w 的是输出到客户端的
}

func main() {
	hander := &Hander{}
	err := http.ListenAndServe(":9090", hander) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
