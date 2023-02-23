package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	fmt.Printf("n1= %T\n", n1)

	var n2 int64 =10
	fmt.Printf("n2 type is %T\nn2 size is %d ", n2, unsafe.Sizeof(n2))
}