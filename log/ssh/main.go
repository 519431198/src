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
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("123")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	addr := fmt.Sprintf("%v:%v", ip, port)

	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("unable to creat ssh conn")
	}
	//sshSession，这个 sshSession 是用来远程执行操作的
	sshSession, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("unable to create ssh session")
	}
	shell := "cd /root;ls -l"
	output, err := sshSession.CombinedOutput(shell)
	if err != nil {
		log.Fatal("error occurred:", err)
	}
	fmt.Println(string(output))
}
