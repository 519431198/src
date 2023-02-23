package main

import "fmt"

func main() {
	//第一种
	var a map[string]string
	a = make(map[string]string, 10)
	a["her01"] = "熏悟空"
	a["her02"] = "居八戒"
	a["her03"] = "如来佛"
	fmt.Println(a)

	//第二种
	cities := make(map[string]string)
	cities["01"] = "杭州"
	cities["02"] = "上海"
	cities["03"] = "北京"
	fmt.Println(cities)

	//第三种
	heroes := map[string]string{
		"heroes01": "熏悟空",
		"heroes02": "居八戒",
		"heroes03": "如来佛",
	}
	fmt.Println(heroes)
}
