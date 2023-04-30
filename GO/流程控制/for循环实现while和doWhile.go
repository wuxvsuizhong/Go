package main

import "fmt"

/*
go中没有while以及do...while语句
可以使用for实现
*/
func main() {
	/*
	   while
	*/
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Printf("第%d行输出\n", i)
		i++
	}

	/*
	   do...while
	*/
	var j int = 1
	for {
		fmt.Printf("第%d行输出\n", j)
		j++
		if j > 10 {
			break
		}
	}
}
