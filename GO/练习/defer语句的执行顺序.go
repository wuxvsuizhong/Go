package main

import (
	"fmt"
)

type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Println(elem)
	return &temp{}
}

func main1() {
	tt := &temp{}
	defer tt.Add(1).Add(2).Add(4).Add(5)
	tt.Add(3)
	//return
}
func work2() {
	fmt.Println("---work2---")
}

func work1() func() {
	fmt.Println("---work1---")
	return work2
}

func main() {
	fmt.Println("---main---")
	defer work1()()
	fmt.Println("---main end---")

	r := test_num()
	fmt.Println("r:", r)

	r = test_num2()
	fmt.Println("r:", r)

	r = test_num3()
	fmt.Println("r:", r)
}

func test_num() (x int) {
	x = 5
	defer func() {
		x += 1
	}()
	return
}

func test_num2() int {
	x := 5
	defer func() {
		x += 1
	}()

	return x
}

func test_num3() (x int) {
	//x = 5
	defer func() {
		x++
	}()

	return 6
}
