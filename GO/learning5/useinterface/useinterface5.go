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
	name string
}

// Start 结构体Phone实现接口的Start
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

// Stop 结构体Phone实现接口的Stop
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

// Phone独有的call方法
func (p Phone) call() {
	fmt.Println("手机打电话...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

// Working Computer结构体绑定Working方法
// Computer结构体的Working方法，接收一个Usb接口的类型变量
func (c Computer) Working(usb Usb) {
	usb.Start()
	//如果usb类型是Phone，因为call方法是Phone独有的，所以需要类型断言判别usb本来类型是Phone，如果usb的本来类型不是Phone，则没有call方法，调用会出错
	if usb, ok := usb.(Phone); ok == true {
		usb.call()
	}
	usb.Stop()
}

/*
Usb2 接口和Usb接口 定义了相同的方法；
这并不矛盾，我们可以认为只要实现了Start()和Stop()方法绑定的结构体，也就实现了 Usb和Usb2两个接口
*/
type Usb2 interface {
	Start()
	Stop()
}

func (c Computer) Working2(usb Usb2) {
	usb.Start()
	usb.Stop()
}

func TestCallInterfaceFuncByStructMethod() {
	getfuncinfo.PrintFuncName()
	c := Computer{}
	p := Phone{}

	c.Working(p)

	ca := Camera{}
	c.Working(ca)
	c.Working2(ca)
}

/*
多态数组，数组的元素类型是某个通用的接口类型I，只要实现了该通用接口的方法，那么这个类型，或者结构体就可以放到这个数组中
*/
func PolymorphicSlice() {
	getfuncinfo.PrintFuncName()
	var usbArr [3]Usb

	usbArr[0] = Phone{"小米"}
	usbArr[1] = Camera{"佳能"}
	usbArr[2] = Phone{"Vivo"}
	fmt.Println(usbArr)

	c := Computer{}
	for _, v := range usbArr {
		c.Working(v)
	}
}
