package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 设置服务器的IP地址和端口号
	server := "192.168.200.12:12001"
	// 连接到服务器
	conn, err := net.DialTimeout("tcp", server, 10*time.Second)
	if err != nil {
		fmt.Println("无法连接到服务器: ", err)
		os.Exit(1)
	}
	// 读取服务器的欢迎信息
	//readResponse(conn)
	// 发送用户名和密码进行身份验证
	sendAuthentication(conn, "oms", "999")
	// 读取异常号码文件
	lines := phoneNum()
	// 执行多个命令
	executeCommands(conn, lines)
	// 关闭连接
	conn.Close()
}

func readResponse(conn net.Conn) {
	// 创建一个缓冲区
	buffer := make([]byte, 1024)
	// 从服务器读取数据
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据错误: ", err)
		os.Exit(1)
	}
	// 解析服务器的响应
	response := string(buffer[:n])
	// 将号码提取出来
	Num := strings.TrimLeft(extractMessage(response, "=86", ","), "=86")
	// 提取所需信息
	message := extractMessage(response, "XSCF_xscf", "(ms)")
	// 打印消息
	if message != "" {
		fmt.Printf("号码:%s,位置信息:%s\n", Num, message)
	}
}

func extractMessage(response string, str1 string, str2 string) string {
	// 在响应中查找消息的起始位置和结束位置
	start := strings.Index(response, str1)
	end := strings.Index(response, str2)
	if start == -1 || end == -1 {
		return ""
	}
	// 提取消息
	message := response[start:end]
	// 返回消息
	return message
}

func sendAuthentication(conn net.Conn, username string, password string) {
	// 发送用户名
	writeCommand(conn, username)
	// 等待服务器响应
	readResponse(conn)
	// 发送密码
	writeCommand(conn, password)
	// 等待服务器响应
	readResponse(conn)
}

func executeCommands(conn net.Conn, commands []string) {
	// 执行多个命令
	for _, command := range commands {
		writeCommand(conn, command)
		// 等待服务器响应
		readResponse(conn)
	}
}

func writeCommand(conn net.Conn, command string) {
	// 发送命令到服务器
	_, err := conn.Write([]byte(command + "\r\n"))
	if err != nil {
		fmt.Println("发送命令错误：", err)
		os.Exit(1)
	}
	// 等待服务器响应
	time.Sleep(1 * time.Second)
}

func phoneNum() (lines []string) {

	// 获取程序所在路径
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	// 将文件名和路径进行拼接
	filePath := filepath.Join(binPath, "phoneNum.csv")
	// 打开文件
	fr, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fr.Close()
	reader := bufio.NewReader(fr)
	for {
		line1, _, err := reader.ReadLine()
		line2 := fmt.Sprintf(":TEST-SVC-SRI:MDN=86%s,TYPE=0;", string(line1))
		if err != nil {
			break
		}
		lines = append(lines, strings.TrimSpace(line2))
	}
	return lines
}
