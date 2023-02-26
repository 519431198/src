package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

// clientConfig 创建连接的配置的结构体
type clientConfig struct {
	Host       string      //ip
	Port       int64       // 端口
	Username   string      //用户名
	Password   string      //密码
	Client     *ssh.Client //ssh client
	LastResult string      //最近一次运行的结果
}

func (cliConf *clientConfig) createClient(host string, port int64, username, password string) {
	//将传入的参数赋值给 ClientConfig 结构体字段
	var client *ssh.Client
	var err error
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password
	cliConf.Port = port
	//一般传入四个参数：user，[]ssh.AuthMethod{ssh.Password(password)}, HostKeyCallback，超时时间
	//获取 ssh 所需要的配置
	config := ssh.ClientConfig{
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)
	//获取client
	/*
		建立链接
		调用 Dial 方法建立 ssh 连接,dial 方法三个参数和两个返回值
		1.第一个为网络类型,例如使用面向连接的 TCP 协议
		2.第二个 addr,为目标机器 ip 地址和端口
		3.第三个 config 为连接生命的配置项
		4.Dial 会返回一个 ssh 连接和错误类型
	*/
	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred:", err)
	}
	//获取到客户端链接
	cliConf.Client = client
}
func (cliConf *clientConfig) RunShell(shell string) string {
	var session *ssh.Session
	var err error
	//获取 session 会话，这个session是用来远程执行操作的
	session, err = cliConf.Client.NewSession()
	if err != nil {
		log.Fatalln("error occurred:", err)
	}
	//执行shell,返回一个 []byte 和一个错误类型
	output, err := session.CombinedOutput(shell)
	if err != nil {
		log.Fatalln("error occurred:", err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}
func main() {
	//对结构体变量进行内存分配
	var cliConf = new(clientConfig)
	//cliConf := new(clientConfig)
	cliConf.createClient("10.0.0.4", 22, "root", "123")
	/*
		可以看到我们这里每次执行一条命令都会创建一条session
		这是因为一条session默认只能执行一条命令
		并且两条命令不可以分开写
		比如：
		cliConf.RunShell("cd /opt")
		cliConf.RunShell("ls")
		这两条命令是无法连续的，下面的ls查看的依旧是~目录
		因此我们可以连着写，使用;分割
	*/
	fmt.Println(cliConf.RunShell("cd /root; ls -l"))
	/*
		total 2448
		-rw-------. 1 root root    1274 Feb 25  2022 anaconda-ks.cfg
		-rw-r--r--  1 root root    1072 Oct 19 18:47 default.conf.1
		-rw-r--r--  1 root root     125 Feb 13 22:25 dump.rdb
		-rw-r--r--  1 root root     326 Feb 25  2022 hostnameIp.sh
		-rw-r--r--  1 root root 2489670 Jan 12 21:09 redis-6.2.8.tar.gz
		-rw-r--r--  1 root root       0 Feb 28  2022 test.txt
	*/
}
