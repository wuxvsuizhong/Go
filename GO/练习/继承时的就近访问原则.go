package main

import "fmt"

/*
当结构体或者匿名结构中有相同的字段或者方法时，编译器采用就近访问，如果需要指定访问某个匿名成员结构体中的某个方法和某个字段，需要通过指明匿名结构体的名字来访问
匿名机构体继承的方式，如果本结构体中没有访问的字段，会自动到继承的父结构体中去找
非匿名结构体继承，如果要访问父结构体中的字段，需要逐层指明结构体名称，也就是写全访问路径
*/
type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A sayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B sayOK", b.Name)
}

func (b *B) hello() {
	fmt.Println("B hello", b.Name)
}

type C struct {
	A
	B
}

type D struct {
	a A
}

func main() {
	var b B
	b.Name = "小明" //给B中的Name赋值，但是不会给A中的Name赋值
	b.hello()     //B hello 小明
	b.SayOk()     //B sayOK 小明

	// 按照就近原则，如果要访问结构体中嵌套的匿名结构体，需要写应用的全路径
	b.A.SayOk() //A sayOk
	b.A.hello() //A sayOk
	// 这里没有获取到Name。因为Name只在B中被赋值了，没有给A赋值

	b.A.Name = "小红" //这里给A中的Name赋值了，和B中的Name互不影响，b中的Name仍然是原来的值
	b.hello()       //B hello 小明
	b.SayOk()       //B sayOK 小明

	b.A.SayOk() //A sayOk 小红
	b.A.hello() //A sayOk 小红

	//var c C
	// 编译报错
	// 当结构体中的子成员有相同的字段时，而结构体没有该成员，那么此时必须要明确结构体名来区分
	//c.Name = "小强"

	var d D
	//fmt.Println(d.Name)  //d.Name undefined 因为D结构体继承了A不是按照匿名结构体的方式继承的，
	// D不是按照匿名方式继承机构体A的，那么如果D中没有Name，不会自动去A中寻找，会直接报错
	fmt.Println(d.a.Name) //通过有名方式继承，如果要访问继承的类中的方法，需要写全路径

}
