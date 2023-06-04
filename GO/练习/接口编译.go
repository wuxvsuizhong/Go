package main

import (
	"fmt"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func main() {
	var w worker = person{name: "张三"}
	fmt.Println(w)                 //{张三 <nil>}
	fmt.Printf("worker类型:%T\n", w) //worker类型:main.person
	//w.work()	//panic,结构体person没有实现worker方法
}
