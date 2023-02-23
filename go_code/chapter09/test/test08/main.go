package main

import "fmt"

func find(a [5]int) {
	var tem int
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				tem = a[j]
				a[j] = a[j+1]
				a[j+1] = tem
			}
		}
	}
	fmt.Printf("maxNum=%v,index=%v\nminNum=%v,index=0", a[len(a)-1], len(a)-1, a[0])
}

func main() {
	var arr = [5]int{45, 77, 34, 89, 11}
	find(arr)
}
