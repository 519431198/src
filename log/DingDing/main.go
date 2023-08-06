package main

import (
	"github.com/wanghuiyt/ding"
	"log"
)

func main() {
	//导入包 go get -u github.com/wanghuiyt/ding
	dingAlert()
}

func dingAlert() {
	d := ding.Webhook{
		AccessToken: "31df59595aecb5a54e0b636938c0840bdbbb8997035ed76277949e858a353a3b",
		Secret:      "SEC70ad42a40d7078a72659de1fcf319ead027a9a78a8d2dead22bb3beda333bd10",
	}
	err := d.SendMessageText("数据库未分表!", "18930547397")
	if err != nil {
		log.Fatalln(err)
	}
}
