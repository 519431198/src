package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	sshClient := sshClient("root", "123", "10.0.0.4")
	defer sshClient.Close()
	filename := filename(sshClient, "/root/22")
	for _, name := range filename {
		fmt.Println(name)
	}

	//scanner := bufio.NewScanner(strings.NewReader(filename))

}

func sshClient(user string, passwd string, url string) (client *ssh.Client) {
	// 配置 SSH 客户端信息
	config := &ssh.ClientConfig{
		User: user, // 远程服务器用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd), // 远程服务器密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 跳过远程服务器的主机密钥校验
	}

	// 拨号连接远程服务器
	client, err := ssh.Dial("tcp", url+":22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	return client
}

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

	// 匹配特定命名规则的文件名
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
