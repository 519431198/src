package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func loadFile(filePath string) []string {
	//打开指定文件夹,获取文件句柄
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//延迟关闭文件句柄,防止内存泄露
	defer file.Close()
	//读取目录下所有文件,ReadDir传入负数表示读取目录下所有文件信息,传入n表示读取前n个文件信息
	//最后将所有文件名保存到字符串切片并返回
	fileInfo, _ := file.ReadDir(-1)
	//创建一个文件名切片
	names := make([]string, 0)
	for _, info := range fileInfo {
		names = append(names, info.Name())
	}
	return names
}

func readRecord(filename string) {
	log.Println(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Println(filename + " error")
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text() // line就是每行文本
		// 对line进行处理
		fmt.Println(line)
	}
}

func main() {
	var path = "/Users/wangyi/log/"
	fileName := loadFile(path)
	time.Sleep(time.Second * 10)
	for _, v := range fileName {
		go readRecord(path + v)
	}
	time.Sleep(time.Second * 15)
}
