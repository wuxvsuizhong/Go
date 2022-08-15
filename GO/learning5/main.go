package main

import(
	"mypro/jsontrans"
	//"mypro/stumanage"
	"mypro/useinterface"
)


func main(){
	jsontrans.TestJsonTrans()
	//stumanage.Start()
	useinterface.TestUseInter()
	useinterface.TestUseInter2()
	useinterface.TestInterface2kind()
	useinterface.TestInterInner2()
	useinterface.TestEmpyInter()
	useinterface.Typeassert(100)
	useinterface.Typeassert("test string")
	useinterface.Typeassert2(100)
	useinterface.Typeassert2("asjdgflkasdf")
	useinterface.Typeassert2(false)
	useinterface.Typeassert2(int64(1000))


}
