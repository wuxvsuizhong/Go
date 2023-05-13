package main

import (
	"fmt"
)

type MethodUtils struct {
	//空字段
}

func (m MethodUtils) Print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main(){
	var m MethodUtils
	m.Print()
}
