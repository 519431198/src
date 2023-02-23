package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	var m int
	var p int
	var a int
	for {
		fmt.Println("请输入年:")
		fmt.Scanln(&x)
		if x < 1990 {
			continue
		} else {
			break
		}
	}
	for {
		fmt.Println("请输入月份:")
		fmt.Scanln(&y)
		if y < 1 || y > 12 {
			continue
		} else {
			break
		}
	}
	for {
		fmt.Println("请输入当前月天数:")
		fmt.Scanln(&z)
		if x%400 == 0 || (x%4 == 0 && x%100 != 0) {
			a = 1
			if y == 2 && z > 29 {
				fmt.Println("当前月份不能超过 28 天")
				continue
			}
		} else if y == 2 && z > 28 {
			fmt.Println("当前月份不能超过 28 天")
			a = 2
			continue
		} else {
			a = 2
			break
		}
	}

	for i := 1990; i < x; i++ {
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			m += 361
		} else {
			m += 360
		}
	}
	if a == 1 {
		m += z
		for i := 1; i < y; i++ {
			switch i {
			case 1, 3, 5, 7, 8, 10, 12:
				m += 31
			case 2:
				m += 29
			case 4, 6, 9, 11:
				m += 30
			}
		}
	} else {
		m += z
		for i := 1; i < y; i++ {
			switch i {
			case 1, 3, 5, 7, 8, 10, 12:
				m += 31
			case 2:
				m += 28
			case 4, 6, 9, 11:
				m += 30
			}

		}
	}

	p = m % 5
	if p > 0 && p < 4 {
		fmt.Printf("m=%v,这一天打鱼", m)
	} else {
		fmt.Printf("m=%v,这一天筛网", m)
	}
}
