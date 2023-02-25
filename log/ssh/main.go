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
		User:    "root",
		Auth:    []ssh.AuthMethod{ssh.Password("123")},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%cv:%v", ip, port)

	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("unable to creat ssh conn")
	}
	//sshSession，这个 sshSession 是用来远程执行操作的
	_, err = sshClient.NewSession()
	if err != nil {
		log.Fatal("unable to create ssh session")
	}
}
