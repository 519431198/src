package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func main() {
	var ip = "10.0.0.4"
	var port = 22
	/*
		添加获取 ssh 所需要的配置
		在进行通信之前需要配置一些用于建立连接的相关参数
		下面的代码段中，我们首先是声明了用户名和密码，连接超时时间设置为10秒钟，addr变量定义了目标机器的IP地址以及端口。
		HostKeyCallback项，我们设置了忽略，这是因为SSH协议为客户端提供了两种安全验证方式，一种是基于口令的安全验证，
		也就是我们常常使用的账号密码形式，另外一种则是基于密钥的安全验证，相较于第一种，这种形式的校验方法极大的提升了安全等级，
		缺点则是时间损耗相对较长。
	*/
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("123")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	addr := fmt.Sprintf("%v:%v", ip, port)
	/*
		建立链接
		调用 Dial 方法建立 ssh 连接,dial 方法三个参数和两个返回值
		1.第一个为网络类型,例如使用面向连接的 TCP 协议
		2.第二个 addr,为目标机器 ip 地址和端口
		3.第三个 config 为连接生命的配置项
		4.Dial 会返回一个 ssh 连接和错误类型
	*/
	clientConn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("unable to creat ssh conn")
	}

	//创建 sshSession 会话，随后可用 sshSession 来与目标机器进行通信.通过 NewSession方法实现该操作
	sshSession, err := clientConn.NewSession()
	if err != nil {
		log.Fatal("unable to create ssh session")
	}
	/*
		每次执行一条命令都会创建一条session
		这是因为一条session默认只能执行一条命令
		并且两条命令不可以分开写
		比如：
		cliConf.RunShell("cd /opt")
		cliConf.RunShell("ls")
		这两条命令是无法连续的，下面的ls查看的依旧是~目录
		因此我们可以连着写，使用;分割
		例如: shell := "cd /root;ls -l"
		此外,该方法不能执行交互式命令,仅仅能执行查询类命令
	*/
	//定义需要执行的操作
	shell := "cd /root;cat test.txt"
	//sshSession.CombinedOutput执行命令后返回的是一个切片,需要进行 string() 转换类型才能查看
	output, err := sshSession.CombinedOutput(shell)
	if err != nil {
		log.Fatal("error occurred:", err)
	}
	fmt.Println(string(output))
}
