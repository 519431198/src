package main

import (
	"fmt"
)

func main() {
	//var score [5]float64
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("input value%d\n", i+1)
	//	fmt.Scan(&score[i])
	//}
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("第%d个元素是%v\n", i, score[i])
	//}
	//var numArr01 [3]int = [3]int{1, 2, 3}
	//fmt.Println("numArr01=", numArr01)
	var heroes [3]string = [...]string{"宋江", "吴用", "铁牛"}
	for index, value := range heroes {
		fmt.Printf("heroes[%d]=%v\n", index, value)
	}
}
