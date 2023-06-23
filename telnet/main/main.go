package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/my/repo/telnet/utils"
)

type client struct {
	server string
	user   string
	passwd string
}

func main() {
	cli := client{
		server: "192.168.200.12:12001",
		user:   "oms",
		passwd: "999",
	}

	// 设置服务器的IP地址和端口号
	// server := "192.168.200.12:12001"
	// 连接到服务器
	conn, err := net.DialTimeout("tcp", cli.server, 10*time.Second)
	if err != nil {
		fmt.Println("无法连接到服务器: ", err)
		os.Exit(1)
	}
	// 读取服务器的欢迎信息
	//readResponse(conn)

	// 发送用户名和密码进行身份验证
	utils.SendAuthentication(conn, cli.user, cli.passwd)
	// 执行多个命令
	utils.ExecuteCommands(conn)
	// 关闭连接
	conn.Close()
}
