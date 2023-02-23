package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n1, i int) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	if i == 1 && n1 == num {
		fmt.Print("第一次猜中,天才")
	} else if (i == 2 || i == 3) && n1 == num {
		fmt.Printf("第%v次猜中,一般", i)
	} else if i >= 4 && i <= 9 && n1 == num {
		fmt.Printf("第%v次猜中,终于猜对了", i)
	}
	if i > 9 {
		fmt.Println("一次没猜中菜鸡")
	}
}

func main() {
	var n1 int
	for i := 0; i < 10; {
		i++
		fmt.Println("请猜一个数字")
		fmt.Scanln(&n1)
		f(n1, i)
	}
}
