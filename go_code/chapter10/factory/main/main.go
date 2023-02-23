package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 88.8)
	fmt.Println(stu.GetScore())
}
