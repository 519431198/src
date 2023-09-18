package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	sort.Ints(numbers)
	fmt.Println(numbers)
}
