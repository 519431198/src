package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 Person = p1

	fmt.Println(p2.Age)
	p2.Name = "tom"
	fmt.Printf("p1.Name=%v p2.Name=%v", p1.Name, p2.Name)
}
