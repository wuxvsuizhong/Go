package main

import "fmt"

var Age int = 10 //全局变量在申明时候赋值是可以的

//var nn int
//nn = 11  //在函数外不允许在单独的行给全局变量赋值

//Name := "aaa"  //报错   这句相当于两条语句 var Name sring; Name = "aaa"
//全局变量不能在外部直接赋值
var Name string = "aaa"

func main() {
	fmt.Println("Name:", Name)
}
