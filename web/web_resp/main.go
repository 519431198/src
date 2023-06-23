package main

import (
	"encoding/json"
	"github.com/my/repo/web/web_mysql/model"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	user := model.User{
		Id:       1,
		UserName: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	resp, _ := json.Marshal(&user)
	w.Write(resp)
}

func reWrite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handle)
	http.HandleFunc("/reWrite", reWrite)
	http.ListenAndServe(":8080", nil)
}
