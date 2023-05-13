package main

import (
	"fmt"
)

type Person struct {
	Name   string            //字符串，初始值””
	Age    int               //整型，初始值0
	Scores []int             //切片类型,初始值[]
	prt    *int              //指针类型 初始值nil
	slice  []int             //切片 初始值[]
	map1   map[string]string //map 初始值map[]
}

/*
如果结构体成员中包含有应用类型如：切片，map,指针，那么它们的初始值都是nil,即还没有分配空间
如果需要使用这样的字段，需要先make，才能使用
*/
func main() {
	var p1 Person
	fmt.Println(p1) //{ 0 [] <nil> [] map[]}

	p1.slice = make([]int, 5) // 结构体中切片需要make后才能使用
	p1.slice[0] = 90

	p1.map1 = make(map[string]string) //结构体中map需要make后才能使用
	p1.map1["Key1"] = "abc~"
	fmt.Println(p1) //{ 0 [] <nil> [90 0 0 0 0] map[Key1:abc~]}
}
