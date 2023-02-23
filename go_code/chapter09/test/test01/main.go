package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成 10 个整数(1-100)保存到数组
	//倒序打印
	//求平均值
	//最大值和最大值的下标
	//查找里面是否有 55
	rand.Seed(time.Now().UnixNano())
	var sj [10]int
	var total int
	for i := 0; i < 10; i++ {
		num := rand.Intn(100) + 1
		sj[i] = num
		total += sj[i]
	}
	var num1 int
	for i := 0; i < 10; i++ {
		for j := 0; j < len(sj)-1; j++ {
			if sj[len(sj)-1-j] > sj[len(sj)-2-j] {
				num1 = sj[len(sj)-1-j]
				sj[len(sj)-1-j] = sj[len(sj)-2-j]
				sj[len(sj)-2-j] = num1
			}
		}
	}
	for i := 0; i < len(sj); i++ {
		if sj[i] == 55 {
			fmt.Printf("数组中有 55 下标为%v", i)
			break
		} else if i == len(sj)-1 {
			fmt.Println("数组中没有 55")
		}
	}
	fmt.Printf("sj的值是%v\n平均数为%v\n", sj, float64(total)/float64(len(sj)))
	fmt.Printf("最大值下标=%v,最小值下标=%v", len(sj)-len(sj), len(sj)-1)
}
