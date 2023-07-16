package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	now := time.Now()
	oldTime := now.AddDate(0, 0, -1)
	logFile := "access_" + oldTime.Format("2006-01-02") + ".log"
	fmt.Println(logFile)
	//d := ding.Webhook{
	//	AccessToken: "31df59595aecb5a54e0b636938c0840bdbbb8997035ed76277949e858a353a3b",
	//	Secret:      "SEC70ad42a40d7078a72659de1fcf319ead027a9a78a8d2dead22bb3beda333bd10",
	//}
	//_ = d.SendMessageText("access_" + oldTime.Format("2006-01-02") + ".log")

}
