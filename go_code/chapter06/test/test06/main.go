package main

import "fmt"

func main() {
	var key1 int = 97
	var key2 int = 90
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", key1)
		key1++
	}
	fmt.Println()
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", key2)
		key2--
	}
}
