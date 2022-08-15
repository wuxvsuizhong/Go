package nonamefunc

import "fmt"

func TestNoname(){
	var ret int = func(n1 int,n2 int) int {
		return n1+n2
	}(100,200)

	fmt.Println("计算结果ret的值:",ret)

	//用变量名承接匿名函数后，可以直接用变量名多次调用匿名函数
	sub_ret := func(n1 int,n2 int) int {
		return n1-n2
	}

	result := sub_ret(100,20)
	fmt.Println("sun_ret计算结果为:",result)
	result = sub_ret(100,30)
	fmt.Println("sun_ret计算结果为:",result)
}

//用全局变量承接匿名函数后，匿名函数可被全局调用
var GnonameFunc = func(n1 int,n2 int) int {
	fmt.Println("全局匿名函数GnonameFunc 被调用...")
	return n1*n2

}

func GetSum() func(int) int {
	var sum int = 0
	return func(n int) int {
		sum += n
		return sum
	}
}
