package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "192.168.31.140:8888")
	if err != nil {
		fmt.Println("Listen() err=", err)
		return
	}
	defer listen.Close()
	for {
		//fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v ip address=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	//循环接收客户端发送数据
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		//等待客户端通过 conn 发送信息
		//如果客户端没有 write[发送],name 协程就阻塞在这里
		//fmt.Println("服务器在等待客户端发送信息", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出 err=", err)
			return
		}
		strTemp := "CofoxServer got msg \"" + string(buf[:n]) + "\" at " + time.Now().String()
		conn.Write([]byte(strTemp))
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
