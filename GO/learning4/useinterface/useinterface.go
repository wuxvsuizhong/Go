package useinterface
import "fmt"

type Person struct{
	Name string
	Age int
}

//(p Person) 把方法绑定到p上，p关联到Person进而关联到结构体Person上,这样就称p是Person结构体的一个接收者
//go的接口，类似于一种动态绑定的方式给struct添加接口
func (p Person)Getname(){
	fmt.Println(p.Name)
}

func (p Person)AddAge1(){
	//go 传递参数是传值，不会改变源数据
	p.Age += 1
}

func (p *Person)AddRealAge1(){
	//传递指针，才是真的改变了源数据
	p.Age++
}

func Testinteface(){
	p1 := Person{"张三",20}
	p1.Getname()
	fmt.Println(p1.Age)
	//20
	p1.AddAge1()
	fmt.Println(p1.Age)
	//20,p1的Age 没有被改变
	p1.AddRealAge1()
	fmt.Println(p1.Age)
}
