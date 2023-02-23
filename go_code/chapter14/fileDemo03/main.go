package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用 ioutil.ReadFile 一次性将文件读取到位

	file := "/Users/wangyi/Desktop/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v",content)  //[65 108 ...  33 10]
	//将[]byte 转换为 string
	fmt.Printf("%v",string(content))
	//没有显示的 open 文件,因此也不需要现实的 Close 文件
	//因为文件的 Open 和 Close 被封装到了 ReadFile
}
