package rename

import "fmt"

func RenameType(){
	type myInt int
	var num1 myInt = 100
	fmt.Printf("num1 的类型是%T,值是%v\n",num1,num1)

	var num2 int
	//num2 = num1
	//会报错，num1自定义为myInt 后，go编译器认为num2和num1不是一种类型

	num2 = int(num1)
	fmt.Printf("num1转换后num2计算结果的类型是%T,值是%v\n",num2,num2)
}

//自定义函数类型
type callback func(int,int)(int)
func myfunc(arg1 int,arg2 int)(int){
	return arg1+arg2
}

func Callmyfunc(calledfunc callback){
	var n1 int
	var n2 int
	fmt.Println("输入第一个数字:")
	fmt.Scanln(&n1)
	fmt.Println("输入第二个数字:")
	fmt.Scanln(&n2)
	fmt.Printf("计算结果:%v\n",calledfunc(n1,n2))
}

func RunCallFunc(){
	Callmyfunc(myfunc)
}


//对返回值命名后,return 的顺序就无所谓了，不用在return语句后加上变量
func Calc(n1 int,n2 int)(sum int,sub int){
	sub = n1 - n2
	sum = n1 + n2
	return
}
