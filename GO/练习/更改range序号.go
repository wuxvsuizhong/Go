package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, n := range nums {
		i = 6
		fmt.Println(i, n)
		sum += n
	}
	fmt.Println(sum)
}

/*
修改range迭代返回的序号i不会有range遍历有任何影响，range每次返回的value值不会受影响
*/
