package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

// ClientConfig 创建连接的配置的结构体
type ClientConfig struct {
	Host       string      //ip
	Port       int64       // 端口
	Username   string      //用户名
	Password   string      //密码
	Client     *ssh.Client //ssh client
	LastResult string      //最近一次运行的结果
}

func (cliConf *ClientConfig) createClient(host string, port int64, username, password string) {
	var client *ssh.Client
	var err error
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password
	cliConf.Port = port
	//一般传入四个参数：user，[]ssh.AuthMethod{ssh.Password(password)}, HostKeyCallback，超时时间，
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
	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred:", err)
	}
	cliConf.Client = client
}
func (cliConf *ClientConfig) RunShell(shell string) string {
	var (
		session *ssh.Session
		err     error
	)
	//获取session，这个session是用来远程执行操作的
	session, err = cliConf.Client.NewSession()
	if err != nil {
		log.Fatalln("error occurred:", err)
	}
	//执行shell
	output, err := session.CombinedOutput(shell)
	if err != nil {
		log.Fatalln("error occurred:", err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}
func main() {
	cliConf := new(ClientConfig)
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
		total 20
		drwxr-xr-x 3 root root 4096 Nov 18 14:05 hadoop
		drwxr-xr-x 3 root root 4096 Nov 18 14:20 hive
		drwxr-xr-x 3 root root 4096 Nov 18 15:07 java
		drwxr-xr-x 3 root root 4096 Nov  4 23:01 kafka
		drwxr-xr-x 3 root root 4096 Nov  4 22:54 zookeeper
	*/
}
