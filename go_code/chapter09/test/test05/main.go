package main

import "fmt"

func main() {
	var arr [5]int
	var num int
	for i := 0; i < len(arr); i++ {
		fmt.Println("input num")
		fmt.Scanln(&num)
		arr[i] = num
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[len(arr)-1-j] > arr[len(arr)-2-j] {
				num = arr[len(arr)-1-j]
				arr[len(arr)-1-j] = arr[len(arr)-2-j]
				arr[len(arr)-2-j] = num
			}
		}
	}
	fmt.Println(arr)
}
