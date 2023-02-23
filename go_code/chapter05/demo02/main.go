package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("大于18岁，得自己负责!")
	} else {
		fmt.Println("年龄不大这次就放过你!")
	}
}