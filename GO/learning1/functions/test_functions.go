package functions

import "fmt"

func MultiArgs(args...int){
	for i:= 0;i<len(args);i++ {
		fmt.Printf("%d\n",args[i])
	}
}

func Callfunc(){
	a := MultiArgs
	//函数可以直接赋值给变量
	fmt.Printf("a的类型是%T,MultiArgs的类型是%T\n",a,MultiArgs)
	//使用赋值的变量直接调用函数
	a(11,22,33,44,55,66)

}

func Totalcalc(args...int)(int){
		var sum int
		for _,val := range args{
			sum += val
		}
		return sum
}

//函数可作为形参
func Callfunc2(num1 int,num2 int,callback func(...int)(int)){
	var ret int
	ret += callback(100,200,300)
	ret += num1
	ret += num2
	fmt.Printf("最终结果是:%d\n",ret)
}
