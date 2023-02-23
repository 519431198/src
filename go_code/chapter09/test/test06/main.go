package main

import (
	"fmt"
)

func main() {
	var arr = [...]string{"张飞", "赵云", "关羽", "黄忠", "马超", "夏侯惇", "夏侯渊", "张辽", "许褚", "于禁", "黄忠"}
	var str string
	fmt.Println("input a hero")
	fmt.Scanln(&str)
	var num int
	for i := range arr {
		if str == arr[i] {
			num++
			fmt.Printf("找到了%v对应下标为%v\n", str, i)
		}
	}
	fmt.Printf("出现了%v次", num)
}
