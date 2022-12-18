package main

import (
	"container/list"
	"fmt"
)

type node struct {
	chars []rune
	//结构体元素是切片类型，那么如果是根据某个节点不同的复制了很多的副本，这些副本只是在”外层”有了各自的结构体的副本
	//但是这些副本的切片元素里的值，仍然还是同一份，这是基于切片的特性，切片是内存中一块数组区域的映射，复制节点的时候只是复制了元素本身
	//如果元素是一种带有指向意味的类型，那么复制不了它指向的值
	//避免办法就是不用切片，用数组就不会存在这种问题
	count int
}

type node2 struct {
	nums  [2]int
	count int
}

func main() {
	n := node{[]rune{'B', 'A', 'A', 'A'}, 0}
	nn := node2{[2]int{1, 2}, 0}
	q := list.List{}
	qq := list.List{}
	for i := 0; i < 10; i++ {
		n.count, nn.count = i, i
		index1 := i % len(n.chars)
		n.chars[index1] += rune(i)

		index2 := i % len(nn.nums)
		nn.nums[index2] += i
		q.PushBack(n)
		qq.PushBack(nn)
	}

	//for e := qq.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//	v := e.Value.(node2)
	//	fmt.Printf("address:%p\n", &v.nums[0])
	//}

	for e := q.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
		v := e.Value.(node)
		fmt.Printf("%s,address:%p\n", string(v.chars), &v.chars[0])
	}
}
