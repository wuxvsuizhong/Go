package main

import "fmt"

// go中取余数记住一个公式
// a%b = a - a/b * b  以前就是获取到a/b的商，然后用商乘以b,最后用 a - 商*b

func main() {
	//注意取余数后，余数的符号
	fmt.Println("10%4 = ", 10%4)     //2
	fmt.Println("-10%4 = ", -10%4)   // -2
	fmt.Println("-10%-4 = ", -10%-4) // -2
	fmt.Println("10%-4 = ", 10%-4)   // 2
}
