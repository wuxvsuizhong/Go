package main

import (
	"fmt"
	"mypro/usechannel"
	"mypro/workpool"
)

func main() {
	usechannel.Start()
	fmt.Println("-------------------")
	usechannel.Start2()
	fmt.Println("-------------------")
	usechannel.Start3()
	fmt.Println("-------------------")
	usechannel.Start4()
	fmt.Println("-------------------")

	fmt.Println("-------------------")
	usechannel.Start5()

	fmt.Println("-------------------")
	usechannel.TwoChanneReadWrite()

	//循环运行线程池
	workpool.Start()
	fmt.Println("__________end__________")

}
