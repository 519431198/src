package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	var slice []int = make([]int, 1)
	fmt.Println(slice)
	copy(slice, a)
	fmt.Println(slice)
}
