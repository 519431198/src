package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	//r1有四个 int,在内存中是连续分布的
	r1 := Rect{}
	fmt.Printf("r1.leftUp.x 地址=%p,r1.leftUp.y 地址=%p,r1.rightDown.X=%p,r1.rightDown.Y=%p\n",
		&r1.leftUp.X, &r1.leftUp.Y, &r1.rightDown.X, &r1.rightDown.Y)

	//r2有两个 *Point 类型,这两个 *Point 类型的本身地址也是连续的,但是它们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 地址=%p,r2.rightDown=%p\n",
		&r2.leftUp, &r2.rightDown)
	//他们指向的地址不一定是连续的,要看系统运行时是如何分配的
	fmt.Printf("r2.leftUp 地址=%p,r2.rightDown=%p",
		r2.leftUp, r2.rightDown)
}
