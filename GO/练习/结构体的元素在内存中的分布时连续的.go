package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct{
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{7, 8}}

	//r1有4个int值，在内存中是连续分布的
	fmt.Printf("r1.leftUp x:%p\t r1.leftUp.y:%p\t r1.rightDown.x:%p\t r1.rightDown.y:%p\n",
	&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)
	// r1.leftUp x:0xc0000141e0         r1.leftUp.y:0xc0000141e8        r1.rightDown.x:0xc0000141f0     r1.rightDown.y:0xc0000141f8

	r2 := Rect2{&Point{1,1},&Point{3,3}}
	// r1有两个Point元素是指针类型，这两个元素的地址是连续的
	fmt.Printf("r2.leftUp:%p\t r2.rightDown:%p\n",&r2.leftUp,&r2.rightDown)
	//元素指向的内存空间不一定是连续的
	fmt.Printf("r2.leftUp.x:%p\t r2.leftUp.y:%p\t r2.rightDown.x:%p\t r2.rightDown.y:%p\n",
	&r2.leftUp.x,&r2.leftUp.y,&r2.rightDown.x,&r2.rightDown.y)
	// r2.leftUp.x:0xc0000160c0         r2.leftUp.y:0xc0000160c8        r2.rightDown.x:0xc0000160d0     r2.rightDown.y:0xc0000160d8
	// leftUp的x和rightDown的x并不连续
}