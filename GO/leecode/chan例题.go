package main

import (
	"errors"
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
	var a chan error
	a = make(chan error, 1)
	// a <- 100

	a <- errors.New("NG")
	//fmt.Println(er
	fmt.Println(a)
}
