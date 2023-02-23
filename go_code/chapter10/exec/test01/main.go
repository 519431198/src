package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	cat1.Name = "小花"
	cat1.Age = 3
	cat1.Color = "blue"
	cat1.Hobby = "eat fish"
	fmt.Println("cat1=", cat1)

	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("Color=", cat1.Color)
	fmt.Println("Hobby=", cat1.Hobby)
}
