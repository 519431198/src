package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("utils() n1=", n1)
}

func num(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
func main() {
	n1 := 10
	test(n1)
	fmt.Println("service() n1", n1)
}
