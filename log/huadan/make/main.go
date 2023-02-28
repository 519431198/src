package main

import (
	"fmt"
	"github.com/my/repo/log/huadan/sftp"
	"time"
	_ "time"
)

func main() {
	//now := time.Now()
	////当天日志文件名头部
	//fileName1 := fmt.Sprintf(now.Format("2006-01-02"))
	now := time.Now()
	fileName1 := fmt.Sprint(now.AddDate(0, 0, -1).Format("2006-01-02"))

	//fileName1 := "download"
	var cliConf = new(sftp.ClientConfig)
	cliConf.CreateClient("10.0.0.4", 22, "root", "123")
	//name := fmt.Sprintln(cliConf.RunShell("cd /root; ls " + fileName1 + "*.log"))
	name := fmt.Sprintln(cliConf.RunShell("cd /root; ls " + fileName1 + "*"))
	var fileName string
	for _, v := range name {
		if v == 10 {
			if fileName == "" {
				continue
			}
			//fmt.Println(string(v))
			cliConf.Download("/root/"+fileName, "/Users/wangyi/fileName")
			fileName = ""
			continue
		}
		fileName += fmt.Sprintf(string(v))
	}
}
