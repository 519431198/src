package main

import "fmt"

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolume() float64 {
	return box.len * box.width * box.height
}

func main() {
	var box = Box{
		len:    10.4,
		width:  20,
		height: 3.4,
	}
	volume := box.getVolume()
	fmt.Printf("%.2f", volume)
}
