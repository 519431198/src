package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan = %v, 本身地址%p", intChan, &intChan)

	//向管道写入数据

	intChan <- 10
	num := 21
	intChan <- num
	fmt.Printf("\nchannel len = %v\ncap = %v\n", len(intChan), cap(intChan))

	//从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("\nchannel len = %v\ncap = %v\n", len(intChan), cap(intChan))
}
