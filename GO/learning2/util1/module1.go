package util1

import "fmt"

var U1_num1 int = 100
var U1_str1 string = "hello , this is util1"

func init() {
	fmt.Println("util1 中的init 被执行...")
	U1_num1 += 50

	fmt.Println(varInSubMod)
}
