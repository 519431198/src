package main

import (
	"fmt"
	"go_code/复习/encapsulation封装/model"
)

func main() {
	var p = model.NewPerson("tom")
	p.SetAge(19)
	p.SetSal(10500)
	fmt.Println(p.Name, p.GetAge(), p.GetSal())
}
