package main

import (
	"fmt"
	"go_code/复习/factory/model"
)

func main() {
	//var P model.Student
	//P.Name = "tom"
	//P.Score = 77.5
	stu := model.NewStudent("tom", 56.7)
	fmt.Println(stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
