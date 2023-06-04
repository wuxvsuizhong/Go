package myfuncs

//testing模块框架用于单元测试
import "testing"

//编写测试用例去测试功能函数
func TestAddUp(t *testing.T) {
	//  调用功能函数
	ret := AddUp(10)
	if ret != 55 { //判断结果是否正确
		t.Fatalf("AddUp(10)执行结果错误，期望值:%v,实际值:%v", 55, ret)
		//	错误时，Fatalf会在打印后直接退出程序
	}

	// 测试结果如果正确，打印日志
	t.Logf("AppUp(10)执行正确!")

}

/*
调用testing去测试功能函数的好处是不用每次都在去main函数中添加对要测试的函数的调用
而是可直接在命令行中进入到编写了测试用例的文件路径下，使用 go test -v 命令自动调用以_test结尾的文件如：xxx_test.go文件运行里面的一个个函数Test开头的函数
而能够在xxx_test.go中被自动运行的函数的要求是函数名需要以Test开头，并且随后跟随的字符串首字母不能是小写的，函数名称形如TestXxxx
testing框架实际是替换了main函数，然后由框架把以Test开头的函数自动加入到一个内置的main函数中去运行
测试时也可以直接使用go test命令，这个如果运行正确，则没有日志，错误时会有日志，如果带上-v参数那么正确和错误的日志都会输出

如果模块下有多个xxx_test.go的测试文件，那么go test会默认执行所有测试函数。
如果要执行测试某个单独的xxx_test.go，那么需要明确指定xxx_test.go文件名，并带上要测试的功能函数的原文件xxx.go，命令为：go test -v sub_test.go sub.go

如果要只是测试单个方法时，go test后带上--test.run参数然后带上要测试的函数名即可，如命令go test -v --test.run TestSubfunc 只测试TestSubFunc这个函数
*/
