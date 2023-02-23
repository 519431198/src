package main

import (
	"fmt"
	"time"
)

func putNum(numChan chan int) {
	for i := 1; i <= 20; i++ {
		numChan <- i
	}
	close(numChan)
}

func res(numChan chan int, resChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 20)
		num, ok := <-numChan
		if !ok {
			break
		}
		flag = true
		var res int
		for i := 1; i <= num; i++ {
			res += i
		}
		if flag {
			resChan <- res
		}
	}
	fmt.Println("数据取完了")
	exitChan <- true
}

func main() {
	var numChan = make(chan int, 2000)
	var resChan = make(chan int, 2000)
	var exitChan = make(chan bool, 8)
	go putNum(numChan)
	for i := 0; i < 8; i++ {
		go res(numChan, resChan, exitChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
	fmt.Println("main主线程退出")
}
