package main

import (
	"fmt"
	"go_code/chapter10/encapsulate/utils"
)

func main() {
	p := utils.NewPerson("Tom")
	p.SetAge(18)
	p.SetSal(20000)
	fmt.Println(p)
}
