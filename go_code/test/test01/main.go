package main

import (
	"fmt"
	"time"
)

func main() {
	//获取前一天零点时间
	now := time.Now()
	//fileName := fmt.Sprintf(now.Format("2006-01-02"))
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t2, _ := time.ParseInLocation("2006-01-02", now.Format("2006-01-02"), time.Local)
	towTimeStamp := t2.AddDate(0, 0, -1).Unix()
	//towTimeStr := time.Unix(towTimeStamp, 0).Format("2006-01-02 15:04:05")
	towTimeStr := time.Unix(towTimeStamp, 0).Format("2006-01-02 15:04:05")
	fmt.Print(towTimeStr)
	fmt.Println()
	oldTime := now.AddDate(0, 0, -1)
	//获取一天前这一刻的时间
	fmt.Println(oldTime.Format("2006-01-02 15:04:05"))
	//获取前一天日期
	fmt.Println(oldTime.Format("2006-01-02"))
}
