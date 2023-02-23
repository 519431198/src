package main

import (
	"fmt"
)

func month(n1 int, n2 int) int {
	var n3 int
	if n1%400 == 0 || (n1%4 == 0 && n1%100 != 0) {
		switch n2 {
		case 1, 3, 5, 7, 8, 10, 12:
			n3 = 31
		case 2:
			n3 = 29
		case 4, 6, 9, 11:
			n3 = 30
		}
	} else {
		switch n2 {
		case 1, 3, 5, 7, 8, 10, 12:
			n3 = 31
		case 2:
			n3 = 28
		case 4, 6, 9, 11:
			n3 = 30
		}
	}
	return n3
}

func main() {
	for {
		var n1 int
		var n2 int
		fmt.Println("输入年份月份,空格分隔")
		fmt.Scanln(&n1, &n2)
		if n2 > 12 || n2 < 1 {
			fmt.Println("月份输入错误,请重新输入!")
			continue
		}
		n3 := month(n1, n2)
		fmt.Printf("输入日期是%v年%v月,当月有%v天\n", n1, n2, n3)
	}
}
