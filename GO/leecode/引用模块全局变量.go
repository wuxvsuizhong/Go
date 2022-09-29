package main

import (
	"fmt"
	"leecode/comm"
)

func main() {
	var count int = 6
	for i := 0; i < count; i++ {
		(*comm.GetNum())++ //模块全局变量量保持其唯一性不会因为引用次数导致副本拷贝
		fmt.Println((*comm.GetNum()))

		comm.GetInneNum()
	}
}
