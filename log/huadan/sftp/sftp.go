package sftp

import (
	"bufio"
	"encoding/json"
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

// ClientConfig 创建连接的配置的结构体
type ClientConfig struct {
	Host       string      //ip
	Port       int64       // 端口
	Username   string      //用户名
	Password   string      //密码
	sshClient  *ssh.Client //ssh client
	sftpClient *sftp.Client
	LastResult string //最近一次运行的结果
}

// CreateClient 获取 sftp 链接
func (cliConf *ClientConfig) CreateClient(host string, port int64, username, password string) {
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

// RunShell 远程执行 ssh 命令
func (cliConf *ClientConfig) RunShell(shell string) string {
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
func (cliConf *ClientConfig) Upload(srcPath, dstPath string) {
	srcFile, err := os.Open(srcPath) //本地路径
	if err != nil {
		log.Fatalln(err)
		return
	}
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
func (cliConf *ClientConfig) Download(srcPath, dstPath string) {
	srcFile, err := cliConf.sftpClient.Open(srcPath) //远程服务器文件路径
	if err != nil {
		log.Fatalln(srcPath, err)
		return
	}
	dstFile, _ := os.Create(dstPath) //本地路径
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		log.Fatalln("error occurred", err)
	}
	fmt.Println(srcPath, "下载文件完毕!")
}

type logType struct {
	ServiceCode          string
	PhoneNumberA         string
	PhoneNumberAAreaCode string
	PhoneNumberAOperator string
	PhoneNumberAProvince string
	PhoneNumberACity     string
	PhoneNumberB         string
	PhoneNumberBAreaCode string
	PhoneNumberBOperator string
	PhoneNumberBProvince string
	PhoneNumberBCity     string
	PhoneNumberX         string
	PhoneNumberXAreaCode string
	PhoneNumberXOperator string
	PhoneNumberXProvince string
	PhoneNumberXCity     string
	PhoneNumberY         string
	PhoneNumberYAreaCode string
	PhoneNumberYOperator string
	PhoneNumberYProvince string
	PhoneNumberYCity     string
	ExtensionNumber      string
	BindingId            string
	CallId               string
	CallTime             string
	RingingTime          string
	StartTime            string
	ReleaseTime          string
	ReleaseDirection     string
	ReleaseCause         string
	CallRecording        string
	RecordingUrl         string
	RecordingMode        string
	CallType             string
	CallResult           string
	TransferPhoneNumber  string
	TransferReason       string
	CallDuration         int
	CustomerId           string
	CustomerName         string
	OpenId               string
	SmsContent           string
	Ability              string
	SmsCount             int
	Charge6              int
	Charge60             int
}

//打开文件,创建并返回文件句柄
func openFile(filePath string) (file *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err =", err)
		return
	}
	return file
}

//打开一个文件,如果不存在则创建
func writeFile(filePath string) (file *os.File) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
	}
	return file
}

//读取文件的每一行
func writeChannel(file *os.File, byteChan chan []byte) {
	lineReader := bufio.NewReader(file)
	//for i := 0; i < 10000; i++ {
	for {
		line1, err := lineReader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		byteChan <- []byte(line1)
	}
	close(byteChan)
}

func readfile(byteChan chan []byte, file2 *os.File, exitChan chan bool) {
	//for i := 0; i < 10000; i++ {
	for {
		line2, ok := <-byteChan
		if !ok {
			break
		}
		//fmt.Println(string(line2))
		var data logType
		err := json.Unmarshal(line2, &data)
		if err != nil {
			fmt.Println("json.Unmarshal err = ", err)
			break
		}
		if data.Charge60 == 0 && data.OpenId != "403258633c1d4b5a9dc5a4e49b313b19" {
			continue
		}
		str := fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", data.PhoneNumberA, data.PhoneNumberX, data.PhoneNumberB, data.StartTime, data.ReleaseTime, data.Charge60)
		writer := bufio.NewWriter(file2)
		writer.WriteString(str)
		writer.Flush()
		fmt.Println(str)
		//fmt.Printf("%v,%v,%v,%v,%v,%v\n", data.PhoneNumberA, data.PhoneNumberX, data.PhoneNumberB, data.StartTime, data.ReleaseTime, data.Charge60)
	}
	//fmt.Println("取不到数据,退出!")
	exitChan <- true
}
