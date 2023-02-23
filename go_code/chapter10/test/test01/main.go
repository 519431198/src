package main

import "fmt"

type Method struct {
}

func (mu Method) JudgeNum(m *[3][3]int) {
	var arr [3][3]int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			arr[j][i] = m[i][j]
		}
	}
	*m = arr
}

func main() {
	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var mu Method
	mu.JudgeNum(&arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
