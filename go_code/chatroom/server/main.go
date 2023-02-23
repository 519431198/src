package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭 conn
	defer conn.Close()
	buf := make([]byte, 8096)
	//循环读取客户端发送的信息
	for {
		fmt.Println("读取客户端发送的数据")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("读取到的 buf=", buf[:4])
	}
}

func main() {
	fmt.Println("服务器在 8889 端口监听...")
	Listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err+", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := Listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept err=", err)
		}
		//一旦连接成功,则启动一个协程和客户端保持通讯
		go process(conn)
	}

}
