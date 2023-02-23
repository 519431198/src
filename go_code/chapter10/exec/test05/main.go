package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` //`json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//1.创建一个 Monster 变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	//2.将 monster 变量序列化为 json 格式子串

	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
