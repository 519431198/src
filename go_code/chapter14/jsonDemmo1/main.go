package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//演示讲结构体, map, 切片进行序列化
	testStruct()
	fmt.Println()
	testMap()
	fmt.Println()
	testSlice()
	fmt.Println()
	testFloat64()
}

//对基本数据类型序列化
func testFloat64() {
	var num1 float64 = 3848.33
	marshal, err := json.Marshal(num1)
	if err != nil {
		return
	}
	fmt.Printf("序列化后: %v", string(marshal))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "baby"
	m2["age"] = "2"
	m2["address"] = "上海"
	slice = append(slice, m2)
	marshal, err := json.Marshal(slice)
	if err != nil {
		return
	}
	fmt.Printf("序列化后: %v", string(marshal))

}

//将 map 进行序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	marshal, err := json.Marshal(a)
	if err != nil {
		return
	}
	fmt.Printf("序列化后: %v", string(marshal))
}

// Monster 将结构体序列化
type Monster struct {
	name     string
	Age      string
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		name:     "牛魔王",
		Age:      "500",
		Birthday: "2022-12-11",
		Sal:      8000.1,
		Skill:    "牛魔拳",
	}
	//将 monster 序列化
	marshal, err := json.Marshal(&monster)
	if err != nil {
		return
	}
	fmt.Printf("序列化后: %v", string(marshal))
}
