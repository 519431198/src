package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
}

func (stu *Stu) Say() {
	fmt.Println("Stu say()")
}

func main() {
	var stu Stu = Stu{}
	var u AInterface = &stu
	u.Say()
	fmt.Println()
}
