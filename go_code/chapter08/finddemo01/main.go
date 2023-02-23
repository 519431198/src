package main

import "fmt"

func main() {
	name := [5]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Scan(&heroName)
	//第一种方式
	for i := 0; i < len(name); i++ {
		if heroName == name[i] {
			fmt.Printf("找到 heroName=%v,下标=%v\n", heroName, i)
			break
		} else if i == len(name)-1 {
			fmt.Printf("没有找到%v这个名字\n", heroName)
		}
	}
	//第二种方式(推荐...)
	index := -1
	for i := 0; i < len(name); i++ {
		if heroName == name[i] {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("找到 heroName=%v,下标=%v", heroName, index)
	}
}
