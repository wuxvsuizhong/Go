package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

/*
接口继承的时候被继承的接口中存在同名的方法
*/

type AA interface {
	test01()
	test02()
}

type BB interface {
	test01()
	test03()
}

type CC interface {
	AA
	BB
}

type fortest struct {
}

func (f fortest) test01() {
	fmt.Println("Test01...")
}

func (f fortest) test02() {
	fmt.Println("Test02...")
}

func (f fortest) test03() {
	fmt.Println("Test03...")
}

func InheritIntefaceHasSameMethod() {
	getfuncinfo.PrintFuncName()
	f := fortest{}
	var ci CC
	ci = f
	ci.test01()
}
