package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

type newanimal interface{
	//newanimal 嵌套了interface runner和eater
	runner
	eater
}

type runner interface{
	run()
}

type eater interface{
	eat(string)
}

type newcat struct{
	name string
}

//无论interface如何嵌套，struct只需实现接口中最终定义的方法即可
func (nc newcat)run(){
	fmt.Printf("%s在跑~\n",nc.name)
}

func (nc newcat)eat(food string){
	fmt.Printf("%s在吃%s\n",nc.name,food)
}

func TestInterInner2(){
	getfuncinfo.PrintFuncName()
	var na1 newanimal
	nc1 := newcat{"蓝猫"}
	na1 = nc1
	na1.run()
	na1.eat("小黄鱼")
}
