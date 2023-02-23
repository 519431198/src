package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/my/repo/go_code/chatroom/common/message"
	"net"
)

//写一个函数,完成登陆
func login(userId int, userPwd string) (err error) {
	//下一个开始定协议
	//fmt.Printf("userId = %d userPwd = %s", userId, userPwd)
	//return nil
	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过 conn 发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个 LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4.将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//5.把 data 赋给 mes.data
	mes.Data = string(data)
	//6.将 mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7.这时,data就是要发送的消息
	//7.1 先把 data 的长度发送给服务器
	//先获取到 data的长度 -> 转换成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	//fmt.Printf("客户端发送消息长度成功=%d,%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	//还需要处理服务器端返回的消息

	return
}
