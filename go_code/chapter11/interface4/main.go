package main

import (
	"fmt"
)

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

func (p Phone) Call() {
	fmt.Println("手机打电话")
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

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果 usb 指向 phone 结构体变量则还需要调用 Call 方法
	//类型断言..
	if p, ok := usb.(Phone); ok {
		p.Call()
	}
	usb.Stop()
}

func main() {
	//定义 usb 接口数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"huawei"}
	usbArr[2] = Camera{"suoni"}

	//遍历 usbArr
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

	//fmt.Println(usbArr)
}
