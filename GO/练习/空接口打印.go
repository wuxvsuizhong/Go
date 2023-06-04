package main

import "fmt"

func main() {
	var (
		a int         = 0
		b int64       = 0
		c interface{} = int(0)
		d interface{} = int64(0)
	)

	println(c == 0)
	println(c == a)
	println(c == b)
	println(d == b)
	println(d == 0)
	println(c == 0)

	fmt.Printf("0的默认类型是:%T\n", 0)
	fmt.Printf("0.0的默认类型是:%T\n", 0.0)
	fmt.Printf("'a'的默认类型是:%T\n", 'a')
	fmt.Printf("'abc'的默认类型是:%T\n", "abc")
}
