package main

import (
	"fmt"
	"sync"
	"time"
)

//思路
//1.编写一个函数来计算各个数的阶乘,并放入到一个 map 中
//2.我们启动多个协程,统计的结果放入 map
//3.map 应该是全局的
var myMap = make(map[int]int, 10)
var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将结果放入 myMap 中
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	//开启多个协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠 10 秒钟
	time.Sleep(time.Second * 10)

	//输出结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
