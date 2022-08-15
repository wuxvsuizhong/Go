package main
import(
	"fmt"
	"mypro/myErrProcess"
	"mypro/datastruct"
)


func main(){
	//err_process()
	datastruct.TestDataStruct()
	//datastruct.TestDatastru2()
	datastruct.TestInit()
	datastruct.TestArrType()
	datastruct.TestDataSend()
	datastruct.TestArr2dir()
	datastruct.TravelArr2dir()
	datastruct.TravelArr2dir2()
	datastruct.TestSlice()
	datastruct.TestSlicemake()
	datastruct.TestSliceAppend()
	datastruct.TestSliceCopy()
	datastruct.TestUseMap()
	datastruct.TestMapAttr()
}



func err_process(){
	//defer + 匿名函数的调用 可以捕获并处理错误
	defer func(){
		//启用recover匿名函数可以捕获错误
		err := recover()
		//若错误不为0
		if err != nil{
			fmt.Println("捕获到错误.")
			fmt.Println("err 是:",err)
		}
	}()

	var num1 int = 100
	var num2 int = 10
	fmt.Printf("%v/%v=%v\n",num1,num2,num1/num2)

	num2 = 0

	fmt.Printf("%v/%v=%v\n",num1,num2,num1/num2)

	myErrProcess.TestMyErr(1000,0)
}


