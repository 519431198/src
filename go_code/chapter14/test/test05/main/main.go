package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var monster = Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}
	str := monster.Store()
	var filePath = "/Users/wangyi/test/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(str)
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	fmt.Println()
	ReStore(filePath)
}

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) Store() string {
	marshal, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("monster 序列化后为 %v", string(marshal))
	return string(marshal)
}

func ReStore(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	var monster Monster
	err = json.Unmarshal([]byte(str), &monster)
	if err != nil {
		return
	}
	fmt.Println(monster)
}
