package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("Phone start work!")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop work!")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("Camera start work!")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop work!")
}

func main() {
	//定义 usb 接口数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"huawei"}
	usbArr[2] = Camera{"suoni"}
	fmt.Println(usbArr)
}
