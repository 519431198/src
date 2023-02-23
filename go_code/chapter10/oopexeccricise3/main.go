package main

import (
	"fmt"
)

type Visitor struct {
	Name string
	Age  int
}

func (visit *Visitor) cost() {
	for {
		if visit.Age > 18 && visit.Age < 150 {
			fmt.Printf("name: %v \nage: %v\ncost: 20", visit.Name, visit.Age)
			break
		} else {
			fmt.Println("免费")
			break
		}
	}
}

func main() {
	var visit Visitor
	for {
		fmt.Println("input you name")
		fmt.Scanln(&visit.Name)
		fmt.Println("input you age")
		fmt.Scanln(&visit.Age)
		visit.cost()
	}
}
