package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//通过 now 可以获取到年月日,时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())
	//格式日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d-%d-%d\n",now.Year(),now.Month(),now.Day(),
		now.Hour(),now.Minute(),now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d-%d-%d",now.Year(),now.Month(),now.Day(),
		now.Hour(),now.Minute(),now.Second())
	fmt.Printf("datestr=%v\n",dateStr)

	//格式化日期时间的另一种方式(time.Format)
	//2006-01-02 15:04:05这个数字是固定的符号可以自由更换
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
}
