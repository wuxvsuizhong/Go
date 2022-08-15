package useinterface
import(
	"fmt"
	"mypro/getfuncinfo"
)

type animal2 interface{
	run()
	eat(string)
}

type cat2 struct{
	name string
}

func (c cat2)run(){
	fmt.Printf("%s在跑~\n",c.name)
}

func (c cat2)eat(food string){
	fmt.Printf("%s在吃%s~\n",c.name,food)
}

type bird2 struct {
	name string
}

func (b *bird2)run(){
	fmt.Printf("%s在飞~\n",b.name)
}

func (b *bird2)eat(food string){
	fmt.Printf("%s在吃%s\n",b.name,food)
}

func TestInterface2kind(){
	getfuncinfo.PrintFuncName()
	var a1 animal2

	c1 := cat2{name:"蓝猫",}
	c2 := &cat2{name:"汤姆",}

	a1 = c1
	a1.run()
	a1.eat("小黄鱼")
	//cat接受者是按照值来传递接收的，但是同样可以接受指针，go在内部已经隐含了指针和取内容的转换
	a1 = c2
	a1.run()
	a1.eat("杰瑞")

	//b1 := bird2{name:"大红",}  //出错，bird的接收是按照指针接收的，不能隐式的转换为取内容
	// bird接收者只定义了指针方式传递接收，所以就必须传递指针了
	b1 := &bird2{name:"大红",}
	a1 = b1
	a1.eat("虫子")
}

//结论：结构体接收者如果定义的是按照值传递，那么也可以接收指针,但是如果只定义了按照指针传递，那么只能接收指针

