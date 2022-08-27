package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%v\n", time.Since(start).Nanoseconds())

	var s, seq string
	seq = " "
	start2 := time.Now()
	for _, v := range os.Args[1:] {
		s += seq + v
	}
	fmt.Println(s)
	fmt.Printf("%v\n", time.Since(start2).Nanoseconds())
}
