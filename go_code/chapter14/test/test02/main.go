package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// /Users/wangyi/test

	//创建一个新文件,写入 5 句"Hello , Garden"
	filePath := "/Users/wangyi/test/test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
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

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	//写入5 句"Hello , Garden
	str := "hi! girl!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
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
