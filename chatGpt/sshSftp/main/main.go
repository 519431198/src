package main

import (
	"bufio"
	"fmt"
	"github.com/my/repo/chatGpt/sshSftp/untils"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// 获取ssh client链接
func sshClient(user string, passwd string, url string, port string) (client *ssh.Client) {
	// 配置 SSH 客户端信息
	config := &ssh.ClientConfig{
		User: user, // 远程服务器用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd), // 远程服务器密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 跳过远程服务器的主机密钥校验
	}

	// 拨号连接远程服务器
	client, err := ssh.Dial("tcp", url+":"+port, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	return client
}

// 获取sftp client链接
func sftpClient(client *ssh.Client) (sftpClient *sftp.Client) {
	// 创建新的 SFTP 客户端
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		panic("Failed to create new sftp client: " + err.Error())
	}
	return sftpClient
}

// 获取符合要求的文件名
func filename(sshClient *ssh.Client, remoteDir string) (fileName []string) {
	// 获取远程服务器指定目录下所有文件
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		os.Exit(1)
	}
	defer session.Close()
	cmd := fmt.Sprintf("ls %s", remoteDir)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		os.Exit(1)
	}

	// 正则匹配特定命名规则的文件名,可以根据需求自定义规则
	re := regexp.MustCompile("2023-02-21.\\d{1}\\.log")
	for _, file := range filepath.SplitList(string(output)) {
		if re.MatchString(file) {
			scanner := bufio.NewScanner(strings.NewReader(file))
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				fileName = append(fileName, scanner.Text())
			}
		}
	}
	return fileName
}

// 下载文件
func download(sftpClient *sftp.Client, remotePath string, localPath string) {
	// 获得远程文件名
	remoteFileName := filepath.Base(remotePath)

	// 打开远程文件
	remoteFile, err := sftpClient.Open(remotePath)
	if err != nil {
		panic("Failed to open remote file: " + err.Error())
	}
	defer remoteFile.Close()

	// 获取远程文件信息
	remoteFileInfo, err := remoteFile.Stat()
	if err != nil {
		panic("Failed to get remote file info: " + err.Error())
	}

	// 获取远程文件大小
	remoteFileSize := remoteFileInfo.Size()

	// 创建本地文件
	localFile, err := os.Create(localPath)
	if err != nil {
		panic("Failed to create local file: " + err.Error())
	}
	defer localFile.Close()

	// 读写缓冲区大小
	const bufferSize = 8192
	var buf [bufferSize]byte
	var written int64

	// 定义开始时间
	lastUpdate := time.Now()
	// 逐块读取远程文件并写入本地文件
	for {
		n, err := remoteFile.Read(buf[:])
		if err != nil && err != io.EOF {
			panic("Failed to read from remote file: " + err.Error())
		}
		if n == 0 {
			break
		}
		_, err = localFile.Write(buf[:n])
		if err != nil {
			panic("Failed to write to local file: " + err.Error())
		}

		// 更新已下载的文件大小并输出下载进度
		written += int64(n)
		// time.Since(lastUpdate).Seconds() > 1.5 --> 和上一次时间作比较,超过 1.5s 则输出一次
		if time.Since(lastUpdate).Seconds() > 1.5 {
			fmt.Printf("Downloaded %.2f%%...\n", float64(written)/float64(remoteFileSize)*100)
			//更新当前时间
			lastUpdate = time.Now()
		}
	}

	// 下载完成，输出提示
	fmt.Printf("Downloaded %s to %s\n", remoteFileName, localPath)
}

func main() {

	err := untils.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 连接远程服务器
	sshClient := sshClient(untils.ConfigData.Ssh.User, untils.ConfigData.Ssh.Passwd, untils.ConfigData.Ssh.Ip, untils.ConfigData.Ssh.Port)
	defer sshClient.Close()

	// 创建 SFTP 客户端,用于文件传输
	sftpClient := sftpClient(sshClient)
	defer sftpClient.Close()

	// 获取远程目录下符合要求的文件名
	remoteDir := untils.ConfigData.Path.Filepath1.RemoteDir1
	localPath := untils.ConfigData.Path.Filepath1.LocalDir1
	filename := filename(sshClient, remoteDir)

	// 下载符合要求的文件并保存到本地路径
	for _, name := range filename {
		//fmt.Println(name)
		download(sftpClient, remoteDir+"/"+name, localPath+"/"+name)
	}
}
