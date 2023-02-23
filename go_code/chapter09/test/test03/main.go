package main

import "fmt"

func sz2(a *[4][4]int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len((*a)[i]); j++ {
			if i == len(*a)-1 || i == 0 {
				(*a)[i][j] = 0
			} else if j == 0 || j == len((*a)[i])-1 {
				(*a)[i][j] = 0
			}
		}
	}
}

func main() {
	//var sz = [4][4]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	var sz [4][4]int
	var num int
	for i := 0; i < len(sz); i++ {
		for j := 0; j < len(sz[i]); j++ {
			fmt.Scan(&num)
			sz[i][j] = num
		}
	}
	var arr = sz
	sz2(&arr)
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Print("\t", v2)
		}
		fmt.Println()
	}

}
