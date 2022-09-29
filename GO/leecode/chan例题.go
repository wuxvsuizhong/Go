package main

import (
	"fmt"
)

/*
func main1() {
	var x chan<- chan error
	var y chan (<-chan error)

	fmt.Println(x == y)
}

*/
func main() {
	var a chan<- chan error

	// a <- 100
	a <- chan error

	fmt.Println(a)
}
