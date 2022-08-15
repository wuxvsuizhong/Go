package datastruct

import "fmt"

func TestSlice(){
	var arr [6]int = [6]int{11,22,33,44,55,66}
	fmt.Println("arr:",arr)
	arrsub := arr[1:3]
	fmt.Println("arr[1:3]:",arrsub)
	fmt.Println("len(arr[1:3]",len(arrsub))
	fmt.Println("cap(arr[1:3]",cap(arrsub))

	//切片的本质是一个结构体
	//切片是对数组中的一段数据的一一映射

	fmt.Printf("arr中下标为1的位置的地址:%p\n",&arr[1])
	fmt.Printf("切片arrsub中下标为0的位置的地址:%p\n",&arrsub[0])
	arrsub[0] = 110
	fmt.Println("arr:",arr)
}

func TestSlicemake(){
	//内置函数make定义一个切片类型[],make(type,len,cap)
	slic1 := make([]int,4,10)
	fmt.Println("slic1:",slic1)
	fmt.Println("切片长度len(slic1):",len(slic1))
	fmt.Println("切片容量cap(slic1):",cap(slic1))
	slic1[0] = 11
	slic1[1] = 22
	fmt.Println("赋值后:slic1:",slic1)

	slic2 := []int{44,55,66}
	fmt.Println("slic2:",slic2)
	fmt.Println("slic2长度为:",len(slic2))
	fmt.Println("slic2容量为:",cap(slic2))
}
