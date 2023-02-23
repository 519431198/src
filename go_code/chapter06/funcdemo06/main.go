package main

import "fmt"

func main(){
	res := func (n1 int,n2 int) int {
		return n1 + n2
	}(10,11)
	fmt.Println("res=",res)
}
