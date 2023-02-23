package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	count:=0.0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)+1
		time.Sleep(time.Duration(1))
		fmt.Println("n=", n)
		count++
		if n==99{
			break
		}
	}
	fmt.Printf("count=%v",count)
}
