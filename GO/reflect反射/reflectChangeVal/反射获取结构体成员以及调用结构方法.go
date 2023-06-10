package reflectChangeVal

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Score  float64 `json:"score"`
	Gender string  `json:"gender"`
}

func (m Monster) Print() {
	fmt.Println("-------start--------")
	fmt.Println(m)
	fmt.Println("--------end--------")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float64, gender string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender
}

func CheckStruct(b interface{}) {
	refType := reflect.TypeOf(b) //获取变量的Type
	refVal := reflect.ValueOf(b) //变量的Value
	refKind := refVal.Kind()     //变量的分类
	if refKind != reflect.Struct {
		fmt.Println("传入的变量不是结构体!是 ", refKind)
		return
	}

	//结构体有几个方法
	methodCnt := refVal.NumMethod()
	fmt.Printf("结构体有%v个方法\n", methodCnt)

	//  获取结构体有几个字段
	itemCnt := refVal.NumField()
	fmt.Printf("结构体有%v个字段\n", itemCnt)

	//  遍历结构体的所有字段
	for i := 0; i < itemCnt; i++ {
		fmt.Printf("第%v个字段值：%v\n", i, refVal.Field(i))
		//获取结构体的tag标签，注意标签只能通过reflect.Type来获取
		tagVal := refType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("第%v个字段，标签为：%v\n", i, tagVal)
		}
	}

	/*
	   通过reflect.Value的Method方法，传入方法的序号获取到结构体的某个方法，在用Call发起调用
<<<<<<< HEAD
	   结构体的方法的顺序是按照函数的ASCII排序的，从偶开始计，所以这里其实是调用的Monster的第二个方法Print()
=======
	   结构体的方法的顺序是按照函数的ASCII排序的，从0开始计，所以这里其实是调用的Monster的第二个方法Print()
>>>>>>> dev
	*/
	refVal.Method(1).Call(nil) //即使不传参数，也要写一个nil掺入

	/*
		手动调用结构体的第一个方法GetSum，并传递参数
		Call方法传递的参数是reflect.Value类型的切片，所以如果要传递参数的话，需要参数类型是reflect.Value，如果是普通的数值或者是字符串等需要传递，需要先用reflect.ValueOf获取其reflect.Value类型，相当于做个类型转换
		Call方法返回的也是reflect.Value类型的切片，是结果值的reflect.Value类型的切片，所以要去到结果的真实值，需要reflect.Value.Int()等对应类型的方法得到值
	*/
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) //通过reflect.ValueOf获取到变量的reflect.Value类型，可以理解为把变量转换成reflect.Value类型
	params = append(params, reflect.ValueOf(20))
	res := refVal.Method(0).Call(params)
	fmt.Println("res=", res[0].Int()) //结果是reflect.Value类型的切片，需要调用正确的类型展缓函数得到值
}

func TestCheckStruct() {
	m := Monster{
		Name:   "悟空",
		Age:    600,
		Score:  100,
		Gender: "male",
	}

	CheckStruct(m)
}
