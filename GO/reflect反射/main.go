package main

/*
反射是在运行时确定类型以及值的
*/
import (
	"fmt"
	"reflect"
	"usereflect/reflectChangeVal"
)

/*
编写一个案例，实现对(基本数类型，interface{},reflect.Value的相互转换）
*/

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量的typr,kind,值
	//  1.获取reflect6.Type
	refType := reflect.TypeOf(b)
	fmt.Println("refType:", refType) //refType: int

	//  2.获取reflect.Value
	refVal := reflect.ValueOf(b)
	fmt.Println("refVal:", refVal) //refVal: 100
	/* 通过反射reflect.ValueOf获取的value之不能直接使用
	   n := 2 + refVal   //虽然打印的refVal是100,但是其类型是reflect.Value类型
	   fmt.Println(n)
	*/
	fmt.Printf("refVal的值：%v,类型：%T\n", refVal, refVal) //refVal的值：100,类型：reflect.Value

	//  2.1 把refVal转换成interface{}
	itr := refVal.Interface()

	//  2.2 把interface{}用断言转换成需要的类型
	num := itr.(int)
	fmt.Printf("num:%v,类型是:%T\n", num, num)
	nums := 2 + num //使用断言转换成真正的int后就可以参与运算了
	fmt.Printf("nums:%v\n", nums)

	// 2.3 如果不使用空接口转换，而是要从reflect.Value直接获取变量的值，需要使用reflect.Value这个（结构体）的类型转换的方法
	refnum := refVal.Int()
	fmt.Printf("refnum的值：%v refVal 的类型:%T\n", refnum, refnum) //refnum的值：100 refVal 的类型:int64
	//refnum := refVal.Float()  //会panic.因为refVal对应的变量是Int类型，不能使用其他的类型转换方法

	/*
		获取变量类别kind的2种方法
	*/
	//3 通过reflect.Value获取到变量对应的类别kind
	valkind1 := refVal.Kind()
	fmt.Println(valkind1) //int

	//4 通过reflect.Type获取到变量对应的类别kind
	valkind2 := refType.Kind()
	fmt.Println(valkind2) //int
}

//对结构体使用反射
type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	//对于自定义的结构体，reflect.Type和reflect.Value的kind并不相同
	refType := reflect.TypeOf(b)
	fmt.Printf("结构体reftype：%v\n", refType) //结构体reftype：main.Student

	refVal := reflect.ValueOf(b)
	fmt.Printf("结构体refVal的值：%v,类型:%T\n", refVal, refVal) //结构体refVal的值：{tom 26},类型:reflect.Value

	refKind := refVal.Kind()
	fmt.Printf("结构体refKind的值:%v,类型：%T\n", refKind, refKind) //结构体refKind的值:struct,类型：reflect.Kind
	refKind2 := refType.Kind()
	fmt.Printf("结构体refKind的值:%v,类型：%T\n", refKind2, refKind2) //结构体refKind的值:struct,类型：reflect.Kind
}

func main() {
	var num int = 100
	reflectTest01(num)
	s := Student{"tom", 26}
	reflectTest02(s)

	reflectChangeVal.TestChangeVal()

	reflectChangeVal.TestCheckStruct()
	reflectChangeVal.TestChangeBytag()

}

/*
kind和变量type的区别是：kind是一个反射中自行对各种数据的类型做了一个分类，Kind的本质是一个常量的集合：
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

有时候，变量的kind和type是相同的，如基础的数据类型：int ,float，Uint等
但是如果变量本来是一个结构体struct，那么这时候type就会是：pkg1.Student, mypkg.Str等，带有包名前缀的，而此时kind是上面的常量里的struct，二者这时候就不相同了
所以对于int, flost，uint等这些基础的类型，使用reflect.Value的kind方法，和reflect.Type获取的结果是一样的;
如果是自定义的结构体，那么反射获取的type是和package挂钩的
但是对于自定义的结构体struct，reflect.Value的Kind方法得到的是struct，而Type得到的是pkg.StructName,如本例子中Type就是main.Student;

反射在变量，interface{}， 和reflect.Value之间的互相转换关系如下：
变量<------>interface{}<--------->reflect.Value

通过反射reflect.Valuel来直接获取变量的值，需要数据类型和人对应的获取值的方法对应，如：x是int类型，
那么应该使用reflect.Value().Int() 而不要使用reflect.Value().xxx()其他的类型获取方法，否则会panic
*/
