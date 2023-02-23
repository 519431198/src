package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//conn, err := net.Dial("tcp", "192.168.31.140:8888")
	//if err != nil {
	//	fmt.Println("client dial err=", err)
	//	return
	//}
	////功能一:客户端可以发送单行数据，然后就退出
	//reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	////从终端读取一行用户输入，并准备发送给服务器
	//line, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("readString err=", err)
	//}
	////再将 line 发送给 服务器
	//n, err := conn.Write([]byte(line))
	//if err != nil {
	//	fmt.Println("conn.Write err=", err)
	//}
	//fmt.Printf("客户端发送了 %d 字节的数据，并退出", n)
	conn, err := net.Dial("tcp", "192.168.31.140:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	var line string
	for {
		line = test(conn)
		if line == "exit" {
			break
		}
	}
}

func test(conn net.Conn) (str string) {
	//功能一:客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	//从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err=", err)
	}
	line = strings.Trim(line, " \r\n")
	if line == "exit" {
		return line
	} else {
		_, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		//fmt.Printf("客户端发送了 %d 字节的数据，并退出\n", n)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read() err=", err)
		return
	}
	fmt.Print(string(buf[:n]))
	return line
}
