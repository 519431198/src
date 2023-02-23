package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var arr [26]byte
	//for i := 0; i < 26; i++ {
	//	arr[i] = 'A' + byte(i)
	//}
	//for i := 0; i < 26; i++ {
	//	fmt.Printf("%c ", arr[i])
	//}
	//fmt.Println()
	//
	//var num [5]int = [...]int{1, 2, 100, 5, 6}
	//var maxnum = num[0]
	//var maxindex = 0
	//for i := 1; i < len(num); i++ {
	//	if maxnum < num[i] {
	//		maxnum = num[i]
	//		maxindex = i
	//	}
	//}
	//fmt.Printf("maxnum=%v maxindex=%v", maxnum, maxindex)
	//fmt.Println()

	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("交换前!=", intArr)

	temp := 0
	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[i]
		intArr[i] = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = temp
	}
	fmt.Println("交换后~!=", intArr)
}
