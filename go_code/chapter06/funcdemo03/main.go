package main

import "fmt"

func f(n int) int {
	if n < 1 || n > 10 {
		fmt.Println("输入错误")
	}
	if n == 10{
		return 1
	} else {
		return (f(n+1) + 1) * 2
	}
}
func main(){
	fmt.Printf("一共有%v个桃子",f(1))
}
