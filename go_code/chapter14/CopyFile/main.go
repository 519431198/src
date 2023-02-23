package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile 编写一个函数,接收两个文件路径 srcFileName, dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}

	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			return
		}
	}(srcFile)
	//通过 srcFile, 获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	//通过 dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			return
		}
	}(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	//注意 Copy 函数是 io 包提供的
	//将/Users/wangyi/test/test.txt 文件拷贝到 /Users/wangyi/test1 目录下

	srcFile := "/Users/wangyi/test/test.txt"
	dstFile := "/Users/wangyi/test1/aaa.txt"

	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println(err)
	}
}
