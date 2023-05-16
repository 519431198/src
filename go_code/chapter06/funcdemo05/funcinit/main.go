package main

import "fmt"

func main() {
	fmt.Println("service()...age", age)
}

func init() {
	fmt.Println("init()...")
}

var age = test()

func test() int {
	fmt.Println("utils()...")
	return 90
}
