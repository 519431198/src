package main

import (
	"fmt"
)

func find(arr [8]int) (int, int) {
	var tem int
	var total int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tem = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tem
			}
		}
	}
	min := arr[0]
	max := arr[len(arr)-1]
	for i := 0; i < len(arr); i++ {
		if i == 0 || i == len(arr)-1 {
			continue
		}
		total += arr[i]
	}
	ave := total / len(arr)
	var a [6]int
	for i := 1; i < len(arr)-1; i++ {
		if ave > arr[i] {
			a[i] = ave - arr[i]
		} else if ave < arr[i] {
			a[i] = arr[i] - ave
		} else {
			a[i] = 0
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if i == 0 || i == len(arr)-1 {
			continue
		}
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				tem = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tem
				tem = a[j]
				a[j] = a[j+1]
				a[j+1] = tem
			}
		}
	}
	fmt.Println(ave)
	fmt.Println(a)
	fmt.Println(arr)
	return min, max
}

func main() {
	var arr = [8]int{88, 79, 84, 75, 82, 90, 93, 77}
	fmt.Println(find(arr))
}
