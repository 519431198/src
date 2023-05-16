package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// /Users/wangyi/utils

	//创建一个新文件,写入 5 句"Hello , Garden"
	filePath := "/Users/wangyi/utils/utils.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	//写入5 句"Hello , Garden
	str := "Hello , Garden\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
