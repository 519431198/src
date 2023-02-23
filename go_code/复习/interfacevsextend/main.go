package main

import "fmt"

func main() {
	var Monkey = LittleMonkey{
		Monkey{
			Name: "二师兄",
		},
	}
	Monkey.climbing()
	Monkey.Flying()
	Monkey.Swimming()
}

type Monkey struct {
	Name string
}

// BirdAble 声明接口
type BirdAble interface {
	Flying()
	Swimming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey
}

// Flying 让 LittleMonkey 实现 BirdAble
func (monkey *Monkey) Flying() {
	fmt.Println(monkey.Name, "学会了飞翔...")
}

func (monkey *Monkey) Swimming() {
	fmt.Println(monkey.Name, "学会了游泳...")
}
