package main

import (
	"fmt"
	"time"
)

// 函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	//defer + recover捕获 panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("utils() 发生错误")
	//	}
	//}()

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
