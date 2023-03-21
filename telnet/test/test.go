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

type client struct {
	server string
	user   string
	passwd string
}

func main() {
	// 初始化参数
	num := 1
	phoneNumChan := make(chan string, 100)
	exitChan := make(chan bool, 4)
	client := client{
		server: "192.168.200.12:12001",
		user:   "oms",
		passwd: "999",
	}
	// 连接到服务器
	conn, err := net.DialTimeout("tcp", client.server, 10*time.Second)
	if err != nil {
		fmt.Println("无法连接到服务器: ", err)
		os.Exit(1)
	}
	// 延时关闭链接
	defer conn.Close()
	// 发送用户名和密码进行身份验证
	if err = sendAuthentication(conn, client.user, client.passwd); err != nil {
		fmt.Println("身份验证失败: ", err)
		os.Exit(1)
	}
	// 获取手机号码
	//if err = collectPhoneNum(phoneNumChan); err != nil {
	//	fmt.Println("获取手机号码失败: ", err)
	//	os.Exit(1)
	//}
	collectPhoneNum(phoneNumChan)
	// 启动协程，执行多个命令
	for i := 0; i < num; i++ {
		go executeCommands(conn, exitChan, phoneNumChan)
	}
	// 等待所有协程执行完毕
	for i := 0; i < num; i++ {
		<-exitChan
		fmt.Println(i)
	}
	close(exitChan)
}

// 读取服务器响应
func readResponse(conn net.Conn) error {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据错误: ", err)
		return err
	}
	// 解析服务器的响应
	response := string(buffer[:n])
	// 提取号码
	num := strings.TrimLeft(extractMessage(response, "=86", ","), "=86")
	// 提取所需信息
	message := extractMessage(response, "XSCF_xscf", "(ms)")
	// 打印消息
	if message != "" {
		fmt.Printf("号码:%s,位置信息:%s\n", num, message)
	}
	return nil
}

// 提取消息
func extractMessage(response string, str1 string, str2 string) string {
	start := strings.Index(response, str1)
	end := strings.Index(response, str2)
	if start == -1 || end == -1 {
		return ""
	}
	return response[start:end]
}

// 发送身份验证信息
func sendAuthentication(conn net.Conn, username string, password string) error {
	// 发送用户名
	if err := writeCommand(conn, username); err != nil {
		return err
	}
	// 等待服务器响应
	if err := readResponse(conn); err != nil {
		return err
	}
	// 发送密码
	if err := writeCommand(conn, password); err != nil {
		return err
	}
	// 等待服务器响应
	if err := readResponse(conn); err != nil {
		return err
	}
	return nil
}

// 执行命令
func executeCommands(conn net.Conn, exitChan chan bool, phoneNumChan chan string) {
	// 执行多个命令
	for {
		command, ok := <-phoneNumChan
		if !ok {
			break
		}
		err := writeCommand(conn, command)
		if err != nil {
			fmt.Println("发送命令错误：", err)
			break
		}
		err = readResponse(conn)
		if err != nil {
			fmt.Println("读取命令响应失败: ", err)
			break
		}
	}
	exitChan <- true
}

// 发送命令
func writeCommand(conn net.Conn, command string) error {
	_, err := conn.Write([]byte(command + "\r\n"))
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}

// 获取手机号码
func collectPhoneNum(phoneNumChan chan string) {

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
		line, _, err := reader.ReadLine()
		lines := fmt.Sprintf(":TEST-SVC-SRI:MDN=86%s,TYPE=0;", string(line))
		if err != nil {
			break
		}
		phoneNumChan <- lines
	}
	close(phoneNumChan)
}
