package main

import "fmt"

func main() {
	var s string = "helo中国"              //byte是另外一种整型
	fmt.Println("[]byte(s):", []byte(s)) //按照字节值输出数字

	var bs = []byte{97, 98, 99} // abc
	fmt.Println("string：", string(bs))
}
