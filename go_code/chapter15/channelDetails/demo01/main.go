package main

import "fmt"

func main() {
	//使用 select 可以解决从管道取数据的阻塞问题

	//1.定义一个管道 10 个 int 数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道,5 个数据
	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法再便利管道时,如果不关闭会阻塞而导致 deadlock
	//实际开发中,不好确定什么时候该关闭该管道
	//使用 select 方式解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从 intChan 读取的数据%s\n", v)
		default:
			fmt.Println("取不到数据了")
			return
		}
	}
}
