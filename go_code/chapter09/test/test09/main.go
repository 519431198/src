package main

import "fmt"

func find(arr [8]int) {
	var total int
	var ave int
	var m int
	var n int
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	ave = total / len(arr)
	for i := 0; i < len(arr); i++ {
		if ave > arr[i] {
			m++
		} else if ave < arr[i] {
			n++
		}
	}
	fmt.Println(total)
	fmt.Printf("平均数为%v大于平均数的个数为%v,小于平均数的个数为%v", ave, n, m)

}

func main() {
	var arr = [8]int{88, 79, 84, 75, 82, 90, 93, 77}
	find(arr)
}
