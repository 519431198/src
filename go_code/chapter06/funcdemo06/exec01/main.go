package main

import (
	"fmt"
	_"strings"
)

//func makeSuffix(suffix string) func (string) string {
//	return func (name string) string {
//		if !strings.HasSuffix(name,suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
func f() func(int) int{
	var n = 10
	return func(x int) int{
		n = n + x
		return n
	}
}

func main(){
	//f := makeSuffix(".jpg")
	//fmt.Println("处理后",f("winter"))
	//fmt.Println("处理后",f("bird"))
	a := f()
	fmt.Println("f(1)=",a(1))
	fmt.Println("f(1)=",a(3))
	fmt.Println("f(1)=",a(5))
}
