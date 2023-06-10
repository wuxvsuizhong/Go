package main

import "fmt"

func main() {
	m := make(map[int]int, 3)
	x := len(m)
	m[1] = m[2]
	y := len(m)
	fmt.Println(x, y)

	for k, v := range m {
		fmt.Println("k,v==>", k, v)
	}
}

/*
只有明确指定了key的，才会有key-value键值对
对不存在的键值对取value,如m[k]得到的是类型的默认值
*/
