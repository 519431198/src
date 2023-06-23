package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "请求 URL 是: %s\n\n", r.URL)
	//fmt.Fprintf(w, "请求内容是: %s\n\n", r.URL.RawQuery)
	//fmt.Fprintf(w, "请求头中所有信息: %s\n", r.Header)
	//fmt.Fprintf(w, "请求头中Accept-Encoding信息: %s\n\n", r.Header["Accept-Encoding"])
	//fmt.Fprintf(w, "请求头中Accept-Encoding信息: %s\n\n", r.Header.Get("Accept-Encoding"))
	//// 获取请求体内容的长度
	//len := r.ContentLength
	//// 创建一个切片
	//body := make([]byte, len)
	//// 将请求体中的内容写入 body
	//r.Body.Read(body)
	//// 在浏览器中显示请求体的内容
	//fmt.Fprintf(w, "请求体中的内容是: %s\n\n", string(body))

	// 解析表单,在调用 r.From 之前必须执行该操作
	//r.ParseForm()
	// 获取请求参数
	//fmt.Fprintf(w, "请求参数有: %s\n\n", r.Form)
	//fmt.Fprintf(w, "请求参数有: %s\n\n", r.PostForm)

	// 通过直接调用
	fmt.Fprintf(w, "get 请求中 user参数: %s\n\n", r.FormValue("user"))
	fmt.Fprintf(w, "post 请求中 username 参数: %s\n\n", r.FormValue("username"))
}

func main() {
	http.HandleFunc("/hello", handle)
	http.ListenAndServe(":8080", nil)
}
