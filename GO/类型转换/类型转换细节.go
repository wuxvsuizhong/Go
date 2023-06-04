package main

import (
	"fmt"
	"strconv"
)

func main1() {

	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	//大范围的数转小范围的数，会有溢出，编译没问题，但是结果不可预料
	fmt.Println("num2:", num2) //num2的值是-97,不符合预期

	/*
	   go不带类型自动转换，需要手动转换类型
	*/

	var d1 int32 = 12
	var d2 int64
	var d3 int8

	//d2 = d1 + 20
	d2 = int64(d1) + 20
	//d3 = d1 + 20
	d3 = int8(d1) + 20
	fmt.Println("d2:", d2, "d3:", d3)

	//127在int8的取值范围类，但是计算结果会有溢出，编译不会报错，但是结果错误
	d3 = int8(d1) + 127
	fmt.Println("d3:", d3) //结果为-117 不符合预期

	//d3 = int8(d1) + 128   //编译直接报错因为128超过了int8的最大值

}

func main2() {
	var num1 int = 99
	var num2 float32 = 23.223
	var b bool = true
	var mychar byte = 'c'
	var str string //空的str

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str的类型是 %T, str=%q\n", str, str) //使用%q可以给字符串使用”包裹起来

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)
}

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var str string

	str = strconv.FormatInt(int64(num1), 10) //第二个参数是进制，10就是转成10机制，2就是转成2进制
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)

	str = strconv.Itoa(int(num2)) //另外一种int转string方式，参数只能int类型
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64) //’f'表示格式，10表示精度，也就是保留多少位的小数，64表示数字为float64类型
	//除了'f'之外，其他格式如下：
	//'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	fmt.Printf("str的类型是 %T, str=%q\n", str, str)
}
