package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/wangyi/Desktop/utils.txt")
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	//fmt.Printf("file=%v",*file)

	//当函数退出时,要及时的关闭 file
	defer func(file *os.File) {
		err := file.Close() //要及时关闭文件句柄,否则会有内存泄漏
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(file)

	//创建一个 *reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Println(str)
	}
	//当函数退出时,要及时的关闭 file
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("Close file error")  //要及时关闭文件句柄,否则会有内存泄漏
	//}
	//fmt.Println("文件读取结束")

}
