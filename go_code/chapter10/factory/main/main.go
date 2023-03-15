package main

import (
	"fmt"

	"github.com/my/repo/go_code/chapter10/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 88.8)
	fmt.Println(stu.GetScore())
}
