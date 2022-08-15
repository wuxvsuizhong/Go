package datastruct

import "fmt"

func TestArrType(){
	var arr1 =[3]int{1,2,3}
	fmt.Printf("arr1 类型:%T\n",arr1)
	//arr1 类型:[3]int
	//数组长度是类型的一部分

	var arr2 = [6]int{11,22,33,44,55,66}
	fmt.Printf("arr2 类型:%T\n",arr2)
	//arr1 类型:[6]int
	//数组的长度也是数组的一部分
}

func TestDataSend(){
	var arr3 = [3]int{3,6,9}
	test1(arr3)
	fmt.Println("in TestDataSend ,arr3:",arr3)
	//arr3 仍然是3,6,9 未改变,go中的数组传递是值传递，而非引用传递,相当于拷贝了一份传送过去

	test2(&arr3)
	fmt.Println("调用test2之后，arr3:",arr3)
}

func test1(arr [3]int){
	arr[0] = 7
	fmt.Println("in test01,",arr)
}

func test2(arr *[3]int){
	(*arr)[0] = 7
	fmt.Println("指针传递,test2 中",*arr)
}
