package usestruct

import (
	"fmt"
	"mypro/getfuncinfo"
)

type Teacher struct {
	Name   string
	Age    int
	Gender string
}

func TestUseStruct() {
	/*结构体的使用方式*/
	getfuncinfo.PrintFuncName()
	//1.先声明，再赋值
	var t1 Teacher
	fmt.Println(t1)
	t1.Name = "张安"
	t1.Age = 30
	t1.Gender = "男"
	fmt.Println(t1)
	//2.直接初始化
	var t2 Teacher = Teacher{"李四", 24, "女"}
	fmt.Println(t2)

	//3.new方法创建
	var t3 *Teacher = new(Teacher)
	(*t3).Name = "王五"
	(*t3).Age = 30
	(*t3).Gender = "女"
	fmt.Println(*t3)

	//go直接隐含了对结构体指针的取值操作，所以可以直接使用指针.val 赋值
	var v4 *Teacher = new(Teacher)
	v4.Name = "老刘"
	v4.Age = 29
	v4.Gender = "女"
	fmt.Println(*v4)

	//直接取地址
	var v5 *Teacher = &Teacher{"老七", 40, "男"}
	//v5.Name = "老七"
	//v5.Age = 40
	//v5.Gender = "男"
	fmt.Println(*v5)
}
