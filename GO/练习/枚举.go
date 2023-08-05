package main

import "fmt"

type MsgType uint64

const (
  var1 = iota
  var2
  var3
)

const (
	LoginMsg MsgType = iota
	LoginRes
	RegisterMsg
	RegisterRes
)

const (
	LoginMsg2 MsgType = iota + 5
	LoginRes2
	RegisterMsg2 = iota + 10
	RegisterRes2
)

func main() {
  fmt.Printf("%d,类型:%T\n",var1,var1)
  fmt.Println(var2)
  fmt.Println(var3)

	fmt.Println("--------------------")
	fmt.Printf("%d,类型:%T\n", LoginMsg, LoginMsg)
	fmt.Println(LoginRes)
	fmt.Println(RegisterMsg)
	fmt.Println(RegisterRes)

	fmt.Println("--------------------")
	fmt.Printf("%d,类型:%T\n",LoginMsg2,LoginMsg2)
	fmt.Println(LoginRes2)
	fmt.Println(RegisterMsg2)
	fmt.Println(RegisterRes2)
}
