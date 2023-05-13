package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

/*
深度说明接口和继承的关系
*/

type Monkey struct {
	Name string
}

// BirdFly 鸟类的飞翔技能
type BirdFly interface {
	fly()
}

// FishSwimming 鱼类的游泳技能
type FishSwimming interface {
	Swimming()
}

func (m Monkey) climbind() {
	//  猴子会攀爬的技能
	fmt.Println(m.Name, "生来会爬树")
}

// LittleMonkey 继承自Monkey
type LittleMonkey struct {
	Monkey
}

// LittleMonkey 实现fly接口
func (l *LittleMonkey) fly() {
	fmt.Println(l.Name, "通过学习，会飞翔")
}

// LittleMonkey 实现swimming 接口
func (l LittleMonkey) Swimming() {
	fmt.Println(l.Name, "通过学习，会游泳")
}

func DiffBwtweenInheritAndInterface() {
	getfuncinfo.PrintFuncName()
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbind()
	monkey.fly()
	monkey.Swimming()
}

/*
总结：实现接口可以在不破坏原来继承的关系的情况下，对结构体（类型)进行功能的扩展
*/
