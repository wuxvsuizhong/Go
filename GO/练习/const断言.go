package main

import "fmt"

func main() {
	const X = 7.0
	var x interface{} = X
	//if y, ok := x.(int); ok {
	if y, ok := x.(float32); ok {
		//const类型断言会失败，值y为0
		fmt.Println("if")
		fmt.Println(y)
	} else {
		fmt.Println("else")
		fmt.Println(int(y))
	}

	const NUM = 100
	//对于const类型，强制转换可以
	fmt.Println("const trans", int(NUM))

}
