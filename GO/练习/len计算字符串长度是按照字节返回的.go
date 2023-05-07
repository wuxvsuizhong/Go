package main

import "fmt"

func main() {
	var s string = "hello中国"       //utf8中，每个汉字3个字节 5 + 3 + 3
	fmt.Println("len(s):", len(s)) //11

	var strs = []rune(s)
	fmt.Println("len(strs):", len(strs)) // 7
}
