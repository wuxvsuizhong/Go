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
	workpool.Start()
	fmt.Println("__________end__________")
}
