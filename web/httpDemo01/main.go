package main

import (
	"fmt"
	"net/http"
)

func handle(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, "hello! \n你的访问路由是:", request.URL.Path)
}

func main() {
	// 创建多路复用器
	mux := http.NewServeMux()
	// 注册路由(适配器)
	mux.HandleFunc("/", handle)
	//http.HandleFunc("/", handle) // 适配器
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("http.ListenAndServe err:", err)
	}
}
