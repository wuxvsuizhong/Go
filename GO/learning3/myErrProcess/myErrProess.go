package myErrProcess
import(
	"fmt"
	"errors"
)
func errProcess(num1 int,num2 int)(err error){
	if num2 == 0{
		//抛出自定义错误
		return errors.New("除数不能为0！")
	}
	return nil
}

func TestMyErr(num1 int ,num2 int){
	err := errProcess(num1,num2)
	if err != nil{
		//panic 内置函数会立即停止当前程序	
		panic(err)

	}
	ret := num1/num2
	num2 = 0
	ret = num1/num2
	fmt.Printf("%v/%v = %v\n",num1,num2,ret)
}
