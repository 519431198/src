package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A sayHello", a.age)
}

type B struct {
	A
	Name string
}

func (b *B) haha() {
	fmt.Println("b haha", b.Name)
}

func main() {
	//var b B
	//b.A.Name = "tom"
	//b.A.age = 19
	//b.A.SayOk()
	//b.A.hello()
	//
	//b.Name = "marry"
	//b.age = 22
	//b.SayOk()
	//b.hello()

	var b = B{
		Name: "lara",
	}
	b.haha()
	b.A.SayOk()
}
