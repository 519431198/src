package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

// clientConfig 创建连接的配置的结构体
type clientConfig struct {
	Host       string      //ip
	Port       int64       // 端口
	Username   string      //用户名
	Password   string      //密码
	sshClient  *ssh.Client //ssh client
	sftpClient *sftp.Client
	LastResult string //最近一次运行的结果
}

func (cliConf *clientConfig) createClient(host string, port int64, username, password string) {
	//将传入的参数赋值给 ClientConfig 结构体字段
	var sshClient *ssh.Client
	var sftpClient *sftp.Client
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
	if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred:", err)
	}
	//获取到ssh客户端链接
	cliConf.sshClient = sshClient
	//获取 sftp 链接
	sftpClient, err = sftp.NewClient(sshClient)
	if err != nil {
		log.Fatalln("error occurred:", err)
	}
	//获取到 sftp 客户端链接
	cliConf.sftpClient = sftpClient
}

func (cliConf *clientConfig) RunShell(shell string) string {
	var session *ssh.Session
	var err error
	//获取 session 会话，这个session是用来远程执行操作的
	session, err = cliConf.sshClient.NewSession()
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

// Upload 上传文件
func (cliConf *clientConfig) Upload(srcPath, dstPath string) {
	srcFile, _ := os.Open(srcPath)                   //本地路径
	dstFile, _ := cliConf.sftpClient.Create(dstPath) //远程服务器路径
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	//buf := make([]byte, 2048)
	buf, _ := ioutil.ReadAll(srcFile)
	for {
		_, err := srcFile.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("error occurred:", err)
			} else {
				break
			}
		}
		//_, _ = dstFile.Write(buf[:n])
		_, _ = dstFile.Write(buf)
	}
	fmt.Println(cliConf.RunShell(fmt.Sprintf("ls -l %s", dstPath)))
}

// Download 下载文件
func (cliConf *clientConfig) Download(srcPath, dstPath string) {
	srcFile, _ := cliConf.sftpClient.Open(srcPath) //远程服务器路径
	dstFile, _ := os.Create(dstPath)               //本地路径
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	_, err := srcFile.WriteTo(dstFile)
	if err != nil {
		log.Fatalln("error occurred", err)
	}
	fmt.Println("下载文件完毕!")
}
func main() {
	//对结构体变量进行内存分配
	var cliConf = new(clientConfig)
	//cliConf := new(clientConfig)
	cliConf.createClient("10.0.0.4", 22, "root", "123")
	//执行 shell 命令
	fmt.Println(cliConf.RunShell("cd /root; ls -lh"))
	//本地上传文件到服务器
	cliConf.Upload("/Users/wangyi/test.txt", "/root/test.txt")
	//从服务器中下载文件
	//cliConf.Download("/root/download.txt", "/Users/wangyi/download.txt")
}
