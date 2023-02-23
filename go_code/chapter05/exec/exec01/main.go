package main

import (
	"fmt"
	"math"
)
func main(){
	var n1 int32 = 10
	var n2 int32 =50
	if n1 + n2 > 50 {
		fmt.Println("hello,world!")
	}

	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20 {
		fmt.Println("和=",(n3+n4))
	}

	var num1 int32 = 10
	var num2 int32 = 5
	if (num1 + num2) % 3 == 0 && (num1 + num2) % 5 ==0 {
		fmt.Println("能被3也能被5整除")
	}

	var year int = 2020
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println(year,"是闰年！")
	}

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	var m = b * b - 4 * a * c

	if m > 0 {
		x1 := ( -b + math.Sqrt(m)) / 2 * a
		x2 := ( -b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v",x1,x2)
	} else if m == 0 {
		x1 := ( -b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v",x1)
	} else {
		fmt.Println("无解...")
	}
}
