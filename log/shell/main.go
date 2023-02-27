package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "/Users/wangyi/go/src/go.sum")
	//获取管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return
	}
	//启动命令
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
		return
	}
	//定义一个 bytes.Buffer 类型变量
	var outputBuf0 bytes.Buffer
	tempOutput := make([]byte, 2048)
	//将管道中的数据写入 tempOutput 切片中
	n, err := stdout.Read(tempOutput)
	if err != nil {
		if err == io.EOF {
			return
		} else {
			log.Fatal(err)
			return
		}
	}
	//将 tempOutput 中数据写入缓冲区
	outputBuf0.Write(tempOutput[:n])
	fmt.Printf("%s\n", outputBuf0.String())
}
