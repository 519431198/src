package main

import (
	"fmt"
)

func main() {
	var x float64
	var z string
	var y float64
	fmt.Println("请输入计算式以空格隔开,如: \"6 / 4; 3 + 2\"")
	fmt.Scanln(&x, &z, &y)
	switch z {
	case "+":
		fmt.Printf("%v + %v = %v", x, y, x+y)
	case "-":
		fmt.Printf("%v - %v = %v", x, y, x-y)
	case "*":
		fmt.Printf("%v * %v = %v", x, y, x*y)
	case "/":
		fmt.Printf("%v / %v = %v", x, y, x/y)
	default:
		fmt.Println("请选择正确的计算方式: \"+ - * /\"")
	}
}
