package usedefer

import "fmt"

func TestDefer(n1 int,n2 int) int {
	defer fmt.Println("n1=",n1)
	defer fmt.Println("n2=",n2)
	//defer 后的语句会暂时压栈，不会立即执行，在函数返回时候会按照先进后出的原则反向执行
	//即便是defer会延迟出栈，但是defer 出栈时不会更新值，只会输出当时压栈的时候当时的现场值
	n1 += 10
	n2 += 20
	var sum int = n1 + n2
	fmt.Println("sum=",sum)
	return sum
}
