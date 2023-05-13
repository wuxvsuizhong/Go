package useinterface

import (
	"fmt"
	"mypro/getfuncinfo"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

// 结构体Phone实现接口的Start
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
// 结构体Phone实现接口的Stop
func (p Phone) Stop(){
	fmt.Println("手机停止工作...")
}


type Camera struct{

}

func (c Camera) Start(){
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop(){
	fmt.Println("相机停止工作...")
}

type Computer struct{

}

// Computer结构体绑定Working方法
// Computer结构体的Working方法，接收一个Usb接口的类型变量
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func TestCallInterfaceFuncByStructMethod(){
	getfuncinfo.PrintFuncName()
	c := Computer{}
	p := Phone{}

	c.Working(p)
}