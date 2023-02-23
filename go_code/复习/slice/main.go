package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{2, 6, 3, 32, 56}
	slice := intArr[1:3]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Printf("%p\n", &intArr[1])
	fmt.Printf("%p\n", &slice[0])
	var slice1 = make([]int, 2, 4)
	slice1[0] = 2
	fmt.Println(slice1)
	//切片的遍历
	var arr [5]int = [...]int{2, 4, 6, 7, 9}
	slice2 := arr[1:4]
	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i])
	}
	for i, v := range slice2 {
		fmt.Println("i=", i, "v=", v)
	}
	//append 对切片进行动态增加
	var slice3 = []int{100, 444, 323}
	slice3 = append(slice3, 400, 6666)
	fmt.Println(slice3)
	slice3 = append(slice3, slice3...)
	fmt.Println(slice3)
}
