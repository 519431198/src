package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b占用的空间", unsafe.Sizeof(b))
}
