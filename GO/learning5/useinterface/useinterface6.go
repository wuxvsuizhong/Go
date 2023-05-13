package useinterface

/*
一个结构体可以实现多个接口
*/

import (
	"fmt"
	"mypro/getfuncinfo"
)

// A接口
type Ainterface interface {
	Say()
}

// B接口
type Binterface interface {
	Hello()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("Stu Say()", s.Name)
}

func (s Stu) Hello() {
	fmt.Println("Stu Hello()", s.Name)
}

func OneStructRealizeMultiInterface() {
	getfuncinfo.PrintFuncName()
	s := Stu{"小明"}
	s.Say()   // A接口
	s.Hello() //B接口

	s2 := Stu{"小强"}
	var a2 Ainterface = s2 // s2实现了Ainterface 所以可以可以赋值给Ainterface 变量a2
	var b2 Binterface = s2 // s2实现了Binterface 所以可以可以赋值给Binterface 变量b2
	a2.Say()               //A接口通过访问自己的方法从而调用到结构体实例的方法
	b2.Hello()             //B接口通过调用自己的方法从而调用到结构体实例的方法
}
