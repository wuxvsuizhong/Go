package useinterface
import(
	"fmt"
	"mypro/getfuncinfo"
)

//定义一个接口
type car interface {
	run()
}

type falali struct {
	brand string
}

//结构体对象实现接口
func (f falali) run() {
	fmt.Printf("%s 在跑，速度70迈~\n",f.brand)
}


type baoshijie struct {
	brand string
}
//结构体对象实现接口
func (b baoshijie) run() {
	fmt.Printf("%s 在跑，速度100迈~\n",b.brand)
}

//访问接口的入口
func drive(c car) {
	c.run()
}


func TestUseInter(){
	getfuncinfo.PrintFuncName()
	b1 := baoshijie{brand:"保时捷"}
	f1 := falali{brand:"法拉利"}
	drive(b1)
	drive(f1)

	b1.run()
	f1.run()
}

//接口是不同的结果对象向外界提供某种通用方法的桥梁
//接口定义值只声明方法名，具体实现取决于结构体对象
//接口需要结构体对象实现
