package nonamefunc

import (
	"fmt"
	"strings"
)

func TestNoname() {
	var ret int = func(n1 int, n2 int) int {
		return n1 + n2
	}(100, 200)

	fmt.Println("计算结果ret的值:", ret)

	//用变量名承接匿名函数后，可以直接用变量名多次调用匿名函数
	sub_ret := func(n1 int, n2 int) int {
		return n1 - n2
	}

	result := sub_ret(100, 20)
	fmt.Println("sun_ret计算结果为:", result)
	result = sub_ret(100, 30)
	fmt.Println("sun_ret计算结果为:", result)
}

//用全局变量承接匿名函数后，匿名函数可被全局调用
var GnonameFunc = func(n1 int, n2 int) int {
	fmt.Println("全局匿名函数GnonameFunc 被调用...")
	return n1 * n2

}

/*闭包
闭包中的变量有保持作用，调用后 sum会意识保持，不会随着调动结束而释放
*/
func GetSum() func(int) int {
	var sum int = 0
	return func(n int) int {
		sum += n
		return sum
	}
}

/*
返回闭包
给另一个后缀suffix比如”.jpg"
返回判断传入的字符串s是否是以suffix结束的，如果是那么返回原样的字符串
如果不是那么加上后缀
*/
func MakeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}
