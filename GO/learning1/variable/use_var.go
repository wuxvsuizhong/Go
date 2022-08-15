package variable

import "fmt"

var num1 int = 10
//模块中变量名称小写的变量无法跨package使用，只能在本package中使用
var Num2 int = 100
//大写的变量名称的变量可以被跨package 使用

func ReadvarInpkg(){
	fmt.Println("在本variable模块中num1的值",num1)
	fmt.Println("在本variable模块中Num2的值",Num2)
}

