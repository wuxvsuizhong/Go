package reflectChangeVal

/*
反射操作任意结构体类型
*/

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestOperateStruct(t *testing.T) {
	var (
		model  *user
		refval reflect.Value
	)
	model = &user{}
	refval = reflect.ValueOf(model)
	t.Log("reflact.ValueOf:", refval.Kind().String()) //类型为指针ptr
	refval = refval.Elem()
	t.Log("reflect.Value.Elem:", refval.Kind().String()) //取Elem后类型为struct
	refval.FieldByName("UserId").SetString("12345678")
	refval.FieldByName("Name").SetString("nickname")
	t.Log("model", model)
}
