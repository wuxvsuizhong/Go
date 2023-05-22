package reflectChangeVal

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) Print() {
	fmt.Println("--------start--------")
	fmt.Println(p)
	fmt.Println("--------end--------")
}

func ChangeValByTag(b interface{}, changetag string, changeVal interface{}) {
	refType := reflect.TypeOf(b)
	fmt.Println("refType:", refType)

	//初步片段是否为指针类型
	refPreType := refType.Kind()
	if refPreType != reflect.Ptr {
		fmt.Println("传递的不是指针类型!是：", refPreType)
		return
	}
	//进一步地，通过指针取Elem类获取到数据实体的类型，判断是否为struct
	refKind := refType.Elem().Kind()
	if refKind != reflect.Struct { //检验传入的参数需要是指针
		fmt.Println("传递的不是结构体类型!是：", refKind)
		return
	}

	refVal := reflect.ValueOf(b)
	itemCnt := refVal.Elem().NumField()
	fmt.Println("itemCnt:", itemCnt)
	for i := 0; i < itemCnt; i++ {
		if refType.Elem().Field(i).Tag.Get("json") == changetag { //循环匹配如果tag和要修改的tag一致则往下走
			itemKind := refVal.Elem().Field(i).Kind() //获取要修改的结构体的字段
			fmt.Println("itemKind:", itemKind)
			if itemKind == reflect.TypeOf(changeVal).Kind() { //tag对应的字段的类型和传入的目标值类型一致的时候才会修改
				refVal.Elem().Field(i).Set(reflect.ValueOf(changeVal)) //传入的是指针，要使用Elem
			}
		}
	}
}

/*
通过tag标签实现高自由度的定制修改结构体元素的值
*/
func TestChangeBytag() {
	p := Person{"张三", 30}
	ChangeValByTag(&p, "name", "李四")
	fmt.Println(p)
	ChangeValByTag(&p, "age", 25)
	fmt.Println(p)
}
