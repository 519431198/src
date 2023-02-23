package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量 type, kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v\nrVal type is %T\n", rVal, rVal)
	iVal := rVal.Interface()
	v := iVal.(int)
	fmt.Println(v)
}

type stu struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	//通过反射获取传入的变量 type, kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//fmt.Printf("rVal = %v\nrVal type is %T\n", rVal, rVal)

	//3.获取变量对应的kind
	//(1)rVal.Kind()
	//(2)rTyp.Kind()

	//kind1 := rVal.Kind()
	//kind2 := rTyp.Kind()
	//fmt.Printf("kind1 =%v kind2=%v\n", kind1, kind2)

	iV := rVal.Interface()
	//fmt.Printf("iV = %v iV=%T\n", &iV, &iV)
	v, ok := iV.(*stu)
	if ok {
		v.Name = "marry"
		fmt.Printf("iV = %v iV=%T", v.Name, v)
	}
}

func main() {
	//先定义一个 int
	var num int = 100
	reflectTest01(num)

	stu := stu{
		Name: "tom",
		Age:  11,
	}
	reflectTest02(&stu)
	fmt.Println(stu)
}
