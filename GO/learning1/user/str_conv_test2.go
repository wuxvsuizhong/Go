package user

import (
	"fmt"
	"strconv"
)

func ConvTest() {
	var n1 int = 10
	var s1 string = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("s1 的类型是%T,值是%q\n", s1, s1)

	var n2 float64 = 3.33
	var s2 string = strconv.FormatFloat(n2, 'f', 9, 64)
	fmt.Printf("s2的类型%T,值%q\n", s2, s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3的类型%T,值%q\n",s3,s3)
}

func Foo() {
	fmt.Println("testtttttttt------------")
}

func StrConv(){
	var s1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b 类型%T,值%v\n",b,b)

	var s2 string = "20"
	var num1 int64
	num1,_  = strconv.ParseInt(s2,10,64)
	fmt.Printf("num1的类型%T,值%v\n",num1,num1)

	var s3 string = "3.14"
	var f1 float64
	f1,_ = strconv.ParseFloat(s3,64)
	fmt.Printf("f1的类型%T,值%v\n",f1,f1)

	var s4 string = "golang"
	var b1 bool
	b1,_ = strconv.ParseBool(s4)
	//字符串类型转换，如果不能有效转换，那么最终转换的值是期望类杏的默认值
	fmt.Printf("b1的类型%T,值%v\n",b1,b1)
	//字符串golang转成bool类型就是无效转换，最终b1的值是默认值false

	var num2 int64
	num2,_  = strconv.ParseInt(s4,10,64)
	fmt.Printf("num2 的类型%T,值%v\n",num2,num2)
	//字符串golang转int64是无效转换，最终num2的值是默认值0
}
