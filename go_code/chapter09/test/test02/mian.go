package main

import "fmt"

func main() {
	//在一个升序的数组中插入一个元素,打印该数组,并且其结果仍然是升序
	//87 76 63 38 24 22 10 9 9 3
	var sz = [...]int{87, 76, 63, 38, 24, 22, 10, 9, 5, 3}
	var sz1 [11]int
	var num int
	for i := 0; i < len(sz1)-1; i++ {
		sz1[i] = sz[i]
	}
	fmt.Println("input a num")
	fmt.Scanln(&num)
	sz1[10] = num
	var num2 int
	for i := 0; i < len(sz1)-1; i++ {
		if sz1[len(sz1)-1-i] > sz1[len(sz1)-2-i] {
			num2 = sz1[len(sz1)-1-i]
			sz1[len(sz1)-1-i] = sz1[len(sz1)-2-i]
			sz1[len(sz1)-2-i] = num2
		}
	}
	fmt.Println("新增元素后的数组为", sz1)
}
