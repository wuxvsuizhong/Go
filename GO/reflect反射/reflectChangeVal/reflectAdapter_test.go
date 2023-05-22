package reflectChangeVal

/*
反射实现适配器，使用适配器用作统一处理接口
*/
import (
	"reflect"
	"testing"
)

func TestReflectAdapter(t *testing.T) {
	call1 := func(v1, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value   //指代要调用的函数func
		inValue  []reflect.Value //fucntion 发起调用时候传递的入参列表
		n        int             //入参args的个数
	)

	//反射传入的变量call是函数类型
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	bridge(call1, 1, 2)          //使用适配器调用func1
	bridge(call2, 1, 2, "Hello") //使用使用适配器调用func2
}
