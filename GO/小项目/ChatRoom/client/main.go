package main

import (
	"chatroom/client/login"
	"fmt"
)

func main() {
	var key int
FORLOOP:
	for {
		fmt.Println("--------------------欢迎登录聊天室--------------------")
		fmt.Println("\t\t\t1.用户登录")
		fmt.Println("\t\t\t2.用户注册")
		fmt.Println("\t\t\t3.退出")
		fmt.Print("请选择:")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录")
			login.Login()
		case 2:
			fmt.Println("注册")
		case 3:
			fmt.Println("退出")
			break FORLOOP
		default:
			fmt.Println("输入有误，请重试!")
		}
	}
}
