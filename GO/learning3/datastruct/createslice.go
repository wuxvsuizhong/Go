package datastruct

import (
	"fmt"
	"mypro/getfuncinfo"
)

func TestSliceAppend() {
	//fmt.Println("func testSliceAppend--------------------")
	getfuncinfo.PrintFuncName()
	var intarr [10]int = [10]int{11, 22, 33, 44, 5, 6, 7, 8, 9, 10}

	var slic1 []int = intarr[1:4]
	fmt.Println("slic1", slic1)
	fmt.Printf("slic1地址是:%p\n", &slic1)
	fmt.Println("intarr:", intarr)

	//切片的append实际上是新开辟了一部分空间，把原来的切片的数据部分拷贝一份，然后再添加进去新的数值，返回的是新切片的空间
	slic2 := append(slic1, 111, 222)
	fmt.Println("slic2", slic2)
	fmt.Printf("slic2的地址是:%p\n", slic2)
	//追加后会映射到切片所属的源数组intarr(只要不超出源数组的长度,那么依然会作用到源数组上)
	fmt.Println("intarr:", intarr)
	//intarr: [11 22 33 44 111 222 7 8 9 10]

	//切片追加切片
	slic3 := []int{1, 2}
	//切片追加切片...是固定写法
	slic2 = append(slic2, slic3...)
	fmt.Println("slic2:", slic2)
	fmt.Printf("slic2的地址是:%p\n", slic2)
	fmt.Println("intarr:", intarr)
	//intarr: [11 22 33 44 111 222 1 2 9 10]
}

func TestSliceCopy() {
	getfuncinfo.PrintFuncName()
	var a []int = []int{11, 22, 33, 44, 55, 66}
	var b []int = make([]int, 10)
	copy(b, a)
	fmt.Println("b[]:", b)

	var c []int = make([]int, 2)
	copy(c, a)
	fmt.Println("如果目标slice容量小于源slice，那么拷贝不会报错，但是拷贝结果是目标slice中涵盖源slice中的部分值")
	fmt.Println("c[]:", c)
}

func TestSliceChange() {
	getfuncinfo.PrintFuncName()
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 10

	fmt.Println("slice", slice)
	fmt.Println("slice2", slice2)
	fmt.Println("arr", arr)
}

func change(slice []int) {
	slice[0] = 100
}

func TestSliceChange2() {
	getfuncinfo.PrintFuncName()
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice", slice)
	change(slice)               //切片是引用传递，所以相当于传递了应用或者指针，在子函数中通过切片应用修改了切片映射的值，会反映到源数据中
	fmt.Println("slice", slice) //slice [100 2 3 4 5]
}

func ChangeStringBySlice() {
	getfuncinfo.PrintFuncName()
	//	字符串string本身是一个byte数组，但是string本身是不可变的类型，不能更改
	fmt.Println("如果要修改string中的某个字符，需要先把string转为byte数组切片，通过切片修改元素值，再转换回string")
	s := "hello"
	fmt.Println("s", s)
	arr1 := []byte(s)
	fmt.Println("arr1", arr1)
	arr1[0] = 'z'
	s = string(arr1)
	fmt.Println("s", s)

	// 字符串有中文时，通过byte数组修改字符串会出现乱码，应为byte数组是按照一个字节一个字节来存放字符串的
	fmt.Println("汉字占据3个字节，单独修改某一个字节会导致结果乱码")
	s2 := "hello,中国"
	fmt.Println("s2", s2)
	arr2 := []byte(s2)
	fmt.Println("arr2:", arr2)
	arr2[6] = 'c'
	s2 = string(arr2)
	fmt.Println("s2", s2)

	fmt.Println("把byte数组换用rune数组，rune数组按照字符串里，兼容汉字，这样转换为rune数组，再修改rune数组的元素值就不乱码")
	s3 := "hello,中国"
	fmt.Println("s3", s3)
	arr3 := []rune(s3)
	fmt.Println("arr3", arr3)
	fmt.Println("len(arr3):", len(arr3))
	arr3[6] = '大'
	s3 = string(arr3)
	fmt.Println("s3", s3)

}
