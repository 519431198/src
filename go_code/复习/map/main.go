package main

import "fmt"

func main() {
	//方式一
	var a map[string]string
	//在使用 map 前需要先 make ,make 的作用就是给 map 分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "浪里白条"
	a["no3"] = "许嵩"
	fmt.Println(a)
	//方式二
	city := make(map[string]string)
	city["no1"] = "北京"
	city["no2"] = "上海"
	city["no3"] = "杭州"
	fmt.Println(city)
	//方式三
	heros := map[string]string{
		"hero1": "宋江",
		"hero2": "火舞",
		"hero3": "葫芦娃",
	}
	heros["hero4"] = "刘备"
	fmt.Println(heros)
}
