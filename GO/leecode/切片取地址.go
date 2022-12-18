package main

import "fmt"

func main() {
	s := []string{"A", "B", "C"}
	t := s[:1]
	fmt.Println(&s[0] == &t[0])
	//切片其实就是数组的"引用"
	u := append(s[:1], s[2:]...) // true
	//所以u其实和s是同一片字符串数组区域
	fmt.Println(&s[0] == &u[0]) //true
}
