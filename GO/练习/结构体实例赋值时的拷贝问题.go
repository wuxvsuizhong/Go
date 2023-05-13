package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 30
	p1.Name = "小明"

	//结构体用另外一个结构体赋值的时候，就拷贝了一份
	var p2 = p1   
	fmt.Println(p2.Name, p2.Age)

	//因为是拷贝，赋值互不干扰
	p2.Name = "小李"
	p2.Age = 25
	fmt.Println("p1",p1)   //p1 {小明 30}
	fmt.Println("p2",p2)	//p2 {小李 25}


}