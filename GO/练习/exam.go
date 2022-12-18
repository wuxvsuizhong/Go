package main

import "fmt"

func main() {
	//不能对map 或者map的key或者value取地址
	/*
		m := map[string]int{"uno": 1}
		p := &m["uno"]
		*p = 2
		fmt.Println(m['uno'])
	*/

	s := []interface{}{}
	s = append(s, []interface{}{1, 2, 3}...)
	fmt.Println(s)
	for index, val := range s {
		fmt.Println(index, val)
	}
}
