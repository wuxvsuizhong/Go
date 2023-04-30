package main

import (
	"fmt"

	"mypro/nonamefunc"
	"mypro/usedefer"
	//import先于全局变量初始化以及main函数
	//被初始化的时候按照顺序会先执行import的package里的init函数
	"mypro/util1"
	"mypro/util2"
)

//特殊函数init 先于init 函数执行
func init() {
	fmt.Println("init 函数执行...")
}

func main() {
	fmt.Println("main 函数执行...")
	fmt.Printf("util1.U1_num1的值是:%v\n", util1.U1_num1)
	fmt.Printf("util2.U2_num1的值是:%v\n", util2.U2_num1)

	//--------------------
	nonamefunc.TestNoname()
	var num1 int = 20
	var num2 int = 30
	fmt.Printf("调用全局匿名函数计算%v和%v的乘积计算结果是:%v\n", num1, num2, nonamefunc.GnonameFunc(num1, num2))

	f1 := nonamefunc.GetSum()
	for i := 1; i < 10; i++ {
		fmt.Printf("闭包第%v次计算累计累加结果:%v\n", i, f1(i))
	}
	//-------------------
	usedefer.TestDefer(20, 30)

	/*
		测试返回闭包
	*/
	f := nonamefunc.MakeSuffix(".jpg")
	fmt.Println(f("abc"))       //添加上后缀.jpg
	fmt.Println(f("photo.jpg")) //保原样的photo.jpg输出

}

func getNum() int {
	pre_set := 100
	fmt.Printf("变量pre_set 值:%v\n", pre_set)
	return pre_set
}

//包中的执行顺序是 全局变量-> init -> 函数
//import的时候会执行 全局变量，然后是init，函数只有在被调用的时候会被执行
//全局变量优先于init 被初始化，如果全局变量调用了其他函数，那么其调用的其他函数会优先于init 被调用用于初始化全局变量
var num int = getNum()
