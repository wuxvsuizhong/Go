package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var ps = []Person{
		{name: "张三", age: 10},
		{name: "李四", age: 11},
		{name: "王五", age: 13},
	}

	m := make(map[string]*Person)
	for _, p := range ps { //p只是作为一个临时的加载迭代元素的变量
		fmt.Println(p.name)
		m[p.name] = &p //这里会有问题，p每次迭代都是一份切片元素的拷贝,这里是给map元素地址是临时变量的地址，不是ps每个元素实际的地址
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("k:%s v:%s\n", k, v.name) //这里m中所有的value都是临时变量p最后一次的值
	}
	fmt.Println("-------------------")
	m2 := make(map[string]Person)
	for _, p2 := range ps {
		m2[p2.name] = p2
	}
	fmt.Println(m)
	for k, v := range m2 {
		fmt.Printf("k:%s v:%s\n", k, v.name)
	}
}
