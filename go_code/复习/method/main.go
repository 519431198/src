package main

import (
	"fmt"
)

func main() {
	person := Person{
		name: "老王",
	}
	person.speak()
}

type Person struct {
	name string
}

func (P Person) speak() {
	fmt.Println(P.name, "是一个好人")
}
