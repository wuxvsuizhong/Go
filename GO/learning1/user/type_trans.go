package user

import (
	"fmt"
)

func TranTest() {
	var n1 int = 100
	var n2 float32 = 10.01
	fmt.Printf("%v,%v\n", n1, n2)
	var n3 float32 = float32(n1)
	fmt.Printf("%f,%v\n", n3, n2)

	var c1 byte = 'a'
	fmt.Println(c1)
	fmt.Printf("%c\n", c1)
	var c2 byte = '5'
	fmt.Println(c2)
	fmt.Printf("%c\n", c2)
	var c3 byte = '('
	fmt.Println(c3 + 10)
	fmt.Printf("%c\n", c3)

}
