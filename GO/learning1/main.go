package main

import (
	"fmt"
	"mypro/user"
	"mypro/variable"
	"mypro/progressctl"
	"mypro/functions"
	//import 包时起别名
	myrename "mypro/rename"
)

func main() {
	s := user.Hello()
	fmt.Printf("%v\n", s)
	user.TranTest()
	user.Strtrans()
	user.ConvTest()
	user.Foo()
	user.StrConv()
	user.PointerTest()
	//fmt.Printf("引用variable 模块的num1的值:%v\n",variable.num1)
	//模块中小写字符开头的变量无法跨package 被调用
	fmt.Printf("main引用variable 模块的Num1的值:%v\n",variable.Num2)
	variable.ReadvarInpkg()

	//user.GetInput()
	//user.GetInput2()
	progressctl.TestFor()
	progressctl.TestFor2()
	progressctl.TestGoto()


	functions.MultiArgs(10,20,30,40,50,60)
	functions.Callfunc()
	functions.Callfunc2(3,4,functions.Totalcalc)
	//rename.RenameType()
	myrename.RunCallFunc()
	sum,sub := myrename.Calc(100,10)
	fmt.Printf("两数和为%v,两数差为%v\n",sum,sub)
}
