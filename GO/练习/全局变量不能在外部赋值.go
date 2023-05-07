package main

import "fmt"

var Age = 10

//全局变量在申明时候赋值是可以的

//Name := "aaa"  //报错   这句相当于两条语句 var Name sring; Name = "aaa"
//全局变量不能在外部直接赋值
var Name string = "aaa"

func main() {
	fmt.Println("Name:", Name)
}
