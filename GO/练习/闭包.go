package main

import (
	"fmt"
)

func called() func() int {
	var num int
	return func() int {
		num += 1
		return num
	}

}

func main() {
	f := called()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
