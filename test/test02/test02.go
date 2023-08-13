package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var num int
	fmt.Print("统计往期(1),统计昨天(2): ")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println("2")
	}
}
