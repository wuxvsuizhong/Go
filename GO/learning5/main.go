package main

import (
	"mypro/jsontrans"
	//"mypro/stumanage"
	"mypro/useinterface"
)

func main() {
	jsontrans.TestJsonTrans()
	//stumanage.Start()
	useinterface.TestUseInter()
	useinterface.TestUseInter2()
	useinterface.TestInterface2kind()
	useinterface.TestInterInner2()
	useinterface.TestEmpyInter()
	useinterface.Typeassert(100)
	useinterface.Typeassert("test string")
	useinterface.Typeassert2(100)
	useinterface.Typeassert2("asjdgflkasdf")
	useinterface.Typeassert2(false)
	useinterface.Typeassert2(int64(1000))
	//在某个结构体中调用失效了接口的其他结构体方法
	useinterface.TestCallInterfaceFuncByStructMethod()
	//一个结构体可以实现多个接口
	useinterface.OneStructRealizeMultiInterface()

	useinterface.InheritIntefaceHasSameMethod()

	useinterface.HeroSort()
	//继承和接口的区别
	useinterface.DiffBwtweenInheritAndInterface()
	//多态数组
	useinterface.PolymorphicSlice()

}
