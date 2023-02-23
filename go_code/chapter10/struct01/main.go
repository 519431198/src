package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
}
