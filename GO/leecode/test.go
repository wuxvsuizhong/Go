package main

import (
	"fmt"
	// "strings"
)

// func called1() (msg string, err error){
// 	var chars []string
// 	strings.
// }

func main() {
	type pos [2]int
	a := pos{4, 5}
	b := pos{4, 5}
	// fmt.Printf("%#v\n", a == b)
	fmt.Println(a == b)
}
