package main

import (
	"container/list"
	"fmt"
	// "strings"
)

// func called1() (msg string, err error){
// 	var chars []string
// 	strings.
// }

func main() {
	type pos [2]int
	a := pos{4, 5}
	b := pos{4, 5}
	// fmt.Printf("%#v\n", a == b)
	fmt.Println(a == b) //true

	num := 5
	s := make([]bool, num)
	fmt.Println(s) //[false false false false false]]

	m1 := make(map[int]bool, num)
	m1[1] = true
	fmt.Println(m1[6]) //false

	//tmpval := 0
	//for i := 0; i < 4; i++ {
	//	fmt.Scanln(&tmpval)
	//}

	const str = "hello,gol"
	fmt.Println("len str", len(str))       //9
	fmt.Println("len str[:]", len(str[:])) //9

	var aa byte = 1 << len(str) / 128
	var bb byte = 1 << len(str[:]) / 128
	fmt.Println("aa", aa) //4
	fmt.Println("bb", bb) //0

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)     //[1 2 3 4 5]
	s1 = s1[:len(s1)-1] // 相当于pop最后一个元素
	fmt.Println(s1)     // [1 2 3 4]

	//var str_list []string
	//tmp_s := ""
	//for {
	//	n, _ := fmt.Scanln(&tmp_s)
	//	//if tmp_s == "" {
	//	fmt.Println("n:", n, "tmp_s:", tmp_s)
	//	//break
	//	//}
	//	str_list = append(str_list, tmp_s)
	//}

	fmt.Println(string(rune('A' + 1))) //'B’
	// go内置的链表数据结构
	l1 := list.List{}
	type node struct {
		s []rune
		n int
	}
	tmp_v := []rune("0000")
	//tmp_v[1] += 1
	for i := 0; i < 10; i++ {
		tmp_v[i%3] += 1
		l1.PushBack(node{tmp_v, i})

		//	虽然有循环，但是tmp_v 是一个切片，它是一片数组的映射，所以始终只有一份，所以无论循环多少次，映射的内容以最后一次为准
		//	要想每次都是单独的切片，那么需要使用切片的拷贝: var tmpCopy []rune; copy(tmpCopy,tmp_v)

		//tmpCopy := []rune("0000")
		//copy(tmpCopy, tmp_v)
		//l1.PushBack(node{tmpCopy, i})
	}

	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		/*
			切片中的值都一样，以最后一次的值为准
			{[52 51 51 48] 0}
			{[52 51 51 48] 1}
			......
			{[52 51 51 48] 8}
			{[52 51 51 48] 9}
		*/
	}

}
