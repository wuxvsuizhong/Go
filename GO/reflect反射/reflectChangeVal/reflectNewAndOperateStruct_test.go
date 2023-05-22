package reflectChangeVal

import (
	"reflect"
	"testing"
)

type user2 struct {
	UserId string
	Name   string
}

func TestCreateAndOperateStruct(t *testing.T) {
	var (
		model   *user2
		reftype reflect.Type
		elem    reflect.Value
	)

	reftype = reflect.TypeOf(model)
	t.Log("reftytpe.Kind().string():", reftype.Kind().String()) //ptr model的类型为一个指针

	reftype = reftype.Elem()
	t.Log("reftytpe.Elem().kind():", reftype.Kind().String()) //struct 因为model指针指向的是一个结构体

	elem = reflect.New(reftype)
	//new 返回一个reflect.Value类型值,该值是一个指针,指向一片新申请的结构体内存区(因为reftype的type是*reflectChangeVal.user2)
	//new 能通过去Elem后的实体数据类型，直接创建其结构体类型的实体内存，这就实现了通过reflrect创建结构体实例
	t.Log("reflect.New:", elem.Kind().String())             //ptr
	t.Log("reflect.New.Elem:", elem.Elem().Kind().String()) //struct
	t.Log("reflect.New.elem.type:", elem.Type().String())   //*reflectChangeVal.user2

	model = elem.Interface().(*user2)                //转换类型为实际的结构体类型
	elem = elem.Elem()                               //通过Elem取数据实体
	elem.FieldByName("UserId").SetString("12345678") //修改数据实体中的值
	elem.FieldByName("Name").SetString("nikkaname")
	t.Log("model model.Name:", model, model.Name) //&{12345678 nikkaname} nikkaname
}

/*
可见反射中，New是通过Type来创建结构体实例的，这很好理解，因为只有type是带有报名pacakge前缀的，这样能够在全局中，给自定义的结构体
做全局的唯一类型标识，然后new才知道要创建的是具体哪个包中哪个结构体类型的数据
*/
