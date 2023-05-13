package useinterface
import(
	"mypro/getfuncinfo"
	"fmt"
)

// 定义一个接口animal
type animal interface{
	move()
	eat(sth string)
}

type cat struct{
	name string
	feet int
}

func (c cat) move(){
	fmt.Printf("%s 在移动...\n",c.name)
}

func (c cat)eat(food string){
	fmt.Printf("%s在吃%s~\n",c.name,food)
}

type bird struct{
	name string
	feet int
}

func (b bird) move(){
	fmt.Printf("%s在飞...\n",b.name)
}

func (b bird) eat(food string){
	fmt.Printf("%s在吃%s~\n",b.name,food)
}

//只要实现了接口interface的结构体对象，那么该结构体对象就可以认为是接口的具体实现对象
//接口对象的类型是动态变化的，其初始时候类型是nil,被赋值后主变为具体赋值的interface.struct的类型
func TestUseInter2(){
	getfuncinfo.PrintFuncName()
	//定义一个接口类型的变量a1
	var a1 animal
	fmt.Printf("%T\n",a1)
	//一开始接口类型的变量是nil类型
	c1 := cat{
		name:"蓝猫",
		feet:4,
	}
	c1.eat("小鱼干")
	//c1 所属于的结构体cat实现了animal interface，所以其实例对象也是animal interface类型的实例对象，可以赋值给interface 实例
	a1 = c1
	fmt.Printf("%T\n",a1)
	//给接口类型赋值后，接口类型的变量类型是useinterface.cat

	b1 := bird{
		name:"愤怒的小鸟",
		feet:2,
	}

	b1.eat("虫子")
	//b1 所属于的结构体bird实现了animal interface，所以其实例对象也是animal interface类型的实例对象，可以赋值给interface 实例
	a1 = b1
	fmt.Printf("%T\n",a1)
	//给接口类型赋值后，接口类型的变量类型是useinterface.bird
}
