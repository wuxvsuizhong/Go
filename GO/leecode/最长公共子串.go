package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	short_s := s1
	long_s := s2
	if len(s1) > len(s2) {
		short_s = s2
		long_s = s1
	}

	for i := 1; i < len(long_s); i++ {
		if i >= len(short_s) {
			for j := 0; j < len(short_s); j++ {

			}
		} else {
			for j := len(short_s) - 1 - i; j < len(short_s); j++ {

			}
		}

	}

}
