package main

import (
	"fmt"
)

func getsum( n1 int,n2 int) int {
	return n1 + n2
}

func myfun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1,num2)
}

func main(){
	var a = getsum
	fmt.Printf("a的类型是%T, getsum的类型是%T\n", a, getsum)
	res := a(10,40)	//<==> res := getsum(10,40)
	fmt.Println("res=",res)

	res2 := myfun(getsum, 50,60)
	fmt.Println("res2=", res2)

	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Printf("num1=%v,num2=%v",num1,num2)
}
