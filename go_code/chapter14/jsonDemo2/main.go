package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unmarshalStrut()
	fmt.Println()
	unmarshalMap()
}

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//将 json 反序列化为结构体
func unmarshalStrut() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2022-12-11\"," +
		"\"Sal\":8000.1,\"Skill\":\"牛魔拳\"}"

	//定义一个 monster 的实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		return
	}

	fmt.Println("反序列化后", monster)
}

func unmarshalMap() {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Address\":\"北京\"}"

	//定义一个 monster 的实例
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		return
	}

	fmt.Println("反序列化后", a)
}
