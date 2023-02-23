package main

import "fmt"

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (This *Monkey) climbing() {
	fmt.Println(This.Name, "会爬树")
}

type LittleMonkey struct {
	Monkey
}

func (This LittleMonkey) Flying() {
	fmt.Println(This.Name, "学会飞翔")
}

func (This LittleMonkey) Swimming() {
	fmt.Println(This.Name, "学会游泳")
}

func main() {
	var monkey = LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
