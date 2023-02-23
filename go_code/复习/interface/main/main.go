package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

// Phone :让 phone 实现 Usb 接口的方法
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机打电话")
}

// Camera :让 camera 实现 Usb 接口的方法
type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type computer struct {
}

//Working :编写一个方法 Working 方法,接受一个 Usb 接口类型变量
//只要是实现了 Usb 接口的结构体或者其他都可以传进 Working 方法中
//所谓实现 Usb 接口,就是指实现了 Usb 接口声明的所有方法
func (c computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"iPhone"}
	usbArr[2] = Camera{"尼康"}
	var computer computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
