package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	fileName := "/Users/wangyi/test/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	fmt.Println("Open file success")
	err = file.Close()
	if err != nil {
		fmt.Println("Close file error")
	}
}
