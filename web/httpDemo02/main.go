package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandle struct {
}

func (m *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "自己创建的 handle", r.URL.Path)
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求", r.URL.Path)
}

func main() {
	myHandle := MyHandle{}
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandle,
		ReadTimeout: 2 * time.Second,
	}
	http.Handle("/myHandle", &myHandle)
	//http.ListenAndServe(":8080", nil)
	server.ListenAndServe()

}
