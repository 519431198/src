package main

import (
	"fmt"
	"math/rand"
	"time"
)

func find(arr *[10]int, lIndex int, rIndex int, v int) {
	if lIndex > rIndex {
		fmt.Println("找不到该元素")
		return
	}
	middle := (lIndex + rIndex) / 2
	if (*arr)[middle] > v {
		find(arr, lIndex, middle-1, v)
	} else if (*arr)[middle] < v {
		find(arr, rIndex, middle+1, v)
	} else {
		fmt.Printf("找到了该元素,下标为%v\n", middle)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i := 0; i < 10; i++ {
		num := rand.Intn(100) + 1
		arr[i] = num
	}
	fmt.Println(arr)
	var tem int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tem = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tem
			}
		}
	}
	fmt.Println(arr)

	find(&arr, 0, len(arr)-1, 90)
}
