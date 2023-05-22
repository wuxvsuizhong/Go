package reflectChangeVal

import (
	"fmt"
	"reflect"
)

func ChangeValByReflcet(b interface{}) {
	refVal := reflect.ValueOf(b)
	fmt.Printf("refVal的kind:%v,type:%v\n", refVal.Kind(), refVal.Type()) //refVal的kind:ptr,type:*int

	//因为传递的是变量的指针，所以需要使用reflect.Value的Elem方法获取到变量实际的实体空间，否则会导致修改无效
	refVal.Elem().SetInt(111)
}

func TestChangeVal() {
	var num int = 10

	ChangeValByReflcet(&num) //修改变量的值，使用指针
	fmt.Println(num)
}
