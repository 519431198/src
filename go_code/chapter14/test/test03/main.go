package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//将/Users/wangyi/test/test.txt 文件内容导入到 test1.txt
	//首先将 test.txt 内容读取到内存
	//将读取到的内容写入到test1.txt
	file1Path := "/Users/wangyi/test/test.txt"
	file2Path := "/Users/wangyi/test/test1.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		//读取文件错误
		fmt.Printf("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v", err)
	}

	file1, err := os.Open(file2Path)
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
	}(file1)

	reader := bufio.NewReader(file1)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Println(str)
	}
}
