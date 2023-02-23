package main

import (
	"fmt"
)

func main(){
	//var key byte
	//fmt.Println("请输入一个字符 a b c d e f g")
	//fmt.Scanf("%c",&key)
	//switch key {
	//case 'a':
	//	fmt.Println("周一")
	//case 'b':
	//	fmt.Println("周二")
	//case 'c':
	//	fmt.Println("周二")
	//default:
	//	fmt.Println("输入有误...")
	//
	//}
	//var age int = 10
	//switch {
	//case age == 10 :
	//	fmt.Println("age == 10")
	//case age == 20 :
	//	fmt.Println("age == 20")
	//default:
	//	fmt.Println("没有匹配到")
	//}

	//var score int = 89
	//switch {
	//case score >= 90:
	//	fmt.Println("成绩优秀")
	//case score >= 70 && score < 90:
	//	fmt.Println("成绩优良")
	//case score >= 60 && score < 70:
	//	fmt.Println("成绩及格")
	//default:
	//	fmt.Println("成绩不及格")
	//}

	var x interface{}
	var y = 10.0
	x=y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x type is :%T", i)
	case int:
		fmt.Printf("x type is int")
	case float64:
		fmt.Printf("x type is float64")
	}
}