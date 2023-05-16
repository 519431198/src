package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,world!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 编写一个函数,每隔一秒输出"hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("utils() hello,world!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
