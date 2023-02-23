package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone start work!")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop work!")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera start work!")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop work!")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var computer Computer
	var Phone Phone
	var Camera Camera

	computer.Working(Phone)
	computer.Working(Camera)
}
