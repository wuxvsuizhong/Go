package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	fmt.Printf("c的地址：%p\n", &c) //绑定结构体的方法使用的是值传递方式，会拷贝一份Circle结构体
	return 3.14 * c.radius * 2
}

func main() {
	var c Circle = Circle{2.0}
	fmt.Printf("面积：%v\n", c.getArea())
	fmt.Printf("主函数中,c的地址:%p\n", &c) //和getArea中获取的结构体的地址不一样，因为是两份单独的数据

	//结构体绑定方法时如果不采用指针，会在方法函数中单独拷贝一份结构体数据，这在一定程度上会消耗内存
	// 所以一般建议结构体绑定方法时，采用结构体的指针

	fmt.Printf("面积2：%v\n", (&c).getArea())
}
