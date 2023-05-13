package getfuncinfo

import(
	"fmt"
	"runtime"
)


func PrintFuncName(){
	//func_name,file,line,ok := runtime.Caller(0)
	func_name,_,_,_ := runtime.Caller(1)
	fmt.Println("________",runtime.FuncForPC(func_name).Name(),"___________")
}

