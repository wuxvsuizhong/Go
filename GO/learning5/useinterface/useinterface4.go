package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

type runner interface {
	run()
}

type eater interface {
	eat(string)
}

//newanimal 嵌套了interface runner和eater，也就是继承了runner和eater接口
type newanimal interface {
	runner
	eater
	sleep()
}

type newcat struct {
	name string
}

//无论interface如何嵌套，struct只需实现接口中最终定义的方法即可
//接口方法的实现需要全部都实现，不能少掉某个接口
func (nc newcat) run() {
	fmt.Printf("%s在跑~\n", nc.name)
}

func (nc newcat) eat(food string) {
	fmt.Printf("%s在吃%s\n", nc.name, food)
}

func (nc newcat) sleep() {
	fmt.Printf("%s在睡觉zzzzz...\n", nc.name)
}

func TestInterInner2() {
	getfuncinfo.PrintFuncName()
	var na1 newanimal
	nc1 := newcat{"蓝猫"}
	na1 = nc1
	//newanimal继承了runner和eater接口，而结构体实例nc1实现了newanimal包含的全部接口，所以nc1可以赋值给newanimal接口类型的变量na1
	// 然后调用newanimal接口的全部方法（包括newanimal继承过来的接口)
	na1.run()
	na1.eat("小黄鱼")
	na1.sleep()
}
