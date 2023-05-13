package main

import "fmt"

func main() {
	var a interface{}
	var f float32 = 1.1
	a = f //空接口可以接受任意的数据类型

	y := a.(float32)
	fmt.Printf("y的类型是:%T\n", y) //float32

	y2 := a.(float64)
	//panic: interface conversion: interface {} is float32, not float64
	//因a代表的变量f本身是float32,所以断言转换为其他类型是会panic的
	fmt.Printf("y2的类型是：%T\n", y2)
}
