package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 10
	reflect01(&num)
	//Elem返回 v 持有的的接口保管的值的 Value 封装,或者 v 持有的指针指向的 Value 封装
	fmt.Println(num)
}

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}
