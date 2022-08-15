package datastruct

import "fmt"

func TestSliceAppend(){
	fmt.Println("func testSliceAppend--------------------")
	var intarr [10]int = [10]int{11,22,33,44,5,6,7,8,9,10}

	var slic1 []int = intarr[1:4]
	fmt.Println("slic1",slic1)
	fmt.Printf("slic1地址是:%p\n",&slic1)
	fmt.Println("intarr:",intarr)

	//切片的append实际上是新开辟了一部分空间，把原来的切片的数据部分拷贝一份，然后再添加进去新的数值，返回的是新切片的空间
	slic2 := append(slic1,111,222)
	fmt.Println("slic2",slic2)
	fmt.Printf("slic2的地址是:%p\n",slic2)
	//追加后会映射到切片所属的源数组intarr(只要不超出源数组的长度,那么依然会作用到源数组上)
	fmt.Println("intarr:",intarr)
	//intarr: [11 22 33 44 111 222 7 8 9 10]

	//切片追加切片
	slic3 := []int{1,2}
	//切片追加切片...是固定写法
	slic2 = append(slic2,slic3...)
	fmt.Println("slic2:",slic2)
	fmt.Printf("slic2的地址是:%p\n",slic2)
	fmt.Println("intarr:",intarr)
	//intarr: [11 22 33 44 111 222 1 2 9 10]
}

func TestSliceCopy(){
	var a[]int = []int{11,22,33,44,55,66}
	var b[]int = make([]int,10)
	copy(b,a)
	fmt.Println("b[]:",b)
}
