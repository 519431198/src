package main

import "fmt"

func main() {
	count2 := 0
	for i := 1; i <= 100; i++ {
		count1 := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count1++
			}
		}
		if count1 <= 2 {
			fmt.Printf("i=%v ", i)
			count2++
		}
		if count2 == 5 {
			fmt.Println()
			count2 = 0
		}
	}
}
