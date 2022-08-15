package datastruct

import (
	"fmt"
	"runtime"
)

func PrintFuncName(){
	//func_name,file,line,ok := runtime.Caller(0)
	func_name,_,_,_ := runtime.Caller(1)
	fmt.Println("________",runtime.FuncForPC(func_name).Name(),"___________")
}


func TestUseMap(){
	PrintFuncName()
	var a map[int]string
	//map需要make创建空间才能使用
	a = make(map[int]string,10)
	a[101] = "teststring1"
	a[102] = "teststring2"
	a[103] = "teststring3"
	a[104] = "string1"
	fmt.Println(a)

	//直接make
	b := make(map[int]string)
	b[1] = "张三"
	b[2] = "李四"
	b[3] = "王五"
	fmt.Println(b)

	//直接初始化
	c := map[int]string{
		11:"一一",
		22:"二二",
		33:"三三",
	}
	c[44] = "四四"
	fmt.Println(c)

	//mao删除
	delete(c,44)
	fmt.Println(c)

	//清空
	//go中没有直接的清空操作，两种方式间接的清空
	//1.for循环遍历delete
	//2.make一个新的覆盖，让旧的被gc垃圾回收

	//查找
	val,flag := c[11]
	//找到了，flag为true，否则为false
	fmt.Println(val)
	fmt.Println(flag)
}

func TestMapAttr(){
	PrintFuncName()
	b := make(map[int]string)
	b[1] = "意义"
	b[2] = "二二"
	
	fmt.Println(len(b))
	for key,val := range b {
		fmt.Printf("key:%v,val:%v\t",key,val)
	}
	fmt.Println()

	//多层map
	a := make(map[string]map[int]string)
	a["班级1"] = make(map[int]string)
	a["班级1"][0] = "1-01"
	a["班级1"][1] = "1-02"
	a["班级1"][2] = "1-03"
	a["班级2"] = make(map[int]string)
	a["班级2"][0] = "2-01"
	a["班级2"][1] = "2-02"
	a["班级2"][2] = "2-03"

	for k1,v1 := range a {
		fmt.Println("班级",k1)
		for k2,v2 := range v1 {
			fmt.Printf("key:%v,编号：%v\t",k2,v2)
		}
		fmt.Println()
	}
}
