package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

//单独一个interface只是关键字而已
//interface{} 才是空接口类型
// 空接口可以看成是一种泛型

func TestEmpyInter(){
	getfuncinfo.PrintFuncName()
	var m1 map[string]interface{}
	m1 = make(map[string]interface{},10)
	m1["name"]="张三"
	m1["age"]=30
	m1["isworker"]=true
	m1["hobby"] = [...]string{"抽烟","喝酒","烫头"}
	fmt.Println(m1)
	fmt.Printf("%T\n",m1)

	show(false)
	show(nil)
	show(m1)
}

//空接口作为函数入参
func show(a interface{}){
	getfuncinfo.PrintFuncName()
	fmt.Printf("type:%T,\tvalue:%v\n",a,a)
}

func Typeassert(a interface{}){
	getfuncinfo.PrintFuncName()
	fmt.Printf("%T\n",a)
	str,ok := a.(string)   //断言传入的参数a是一个字符串
	if !ok{
		fmt.Println("不是字符串，猜测错误!\n")
	}else{
		fmt.Println("传递的是一个字符串:",str)
	}
}

func Typeassert2(a interface{}){
	getfuncinfo.PrintFuncName()
	fmt.Printf("传递的是%T\n",a)
	switch t := a.(type){
	case string:
		fmt.Println("传递的是一个字符串:",t)
	case int:
		fmt.Println("传递的是int:",t)
	case int64:
		fmt.Println("传递的是int64:",t)
	case bool:
		fmt.Println("传递的是bool:",t)
	}
}
