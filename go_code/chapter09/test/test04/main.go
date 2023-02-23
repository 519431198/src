package main

import "fmt"

func a(sz *[4][4]int) {
	//var sz1 [4]int
	//sz1 = (*sz)[0]
	//(*sz)[0] = sz[3]
	//(*sz)[3] = sz1
	//sz1 = (*sz)[1]
	//(*sz)[1] = sz[2]
	//(*sz)[2] = sz1
	for i := 0; i < len(*sz)/2; i++ {
		var arr [4]int
		arr = (*sz)[i]
		(*sz)[i] = (*sz)[len(*sz)-1-i]
		(*sz)[len(*sz)-1-i] = arr
	}
}

func main() {
	var sz = [4][4]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	//var sz [4][4]int
	//var num int
	//for i := 0; i < len(sz); i++ {
	//	for j := 0; j < len(sz[i]); j++ {
	//		fmt.Scan(&num)
	//		sz[i][j] = num
	//	}
	//}
	var arr = sz
	a(&arr)
	fmt.Println(arr)
	fmt.Println(arr[0])

}
