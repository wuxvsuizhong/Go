package main

import "fmt"

func main() {
	m := make(map[int]int, 3)
	x := len(m)
	m[1] = m[1]
	y := len(m)
	fmt.Println(x, y)

	for k, v := range m {
		fmt.Println("k,v==>", k, v)
	}
}
