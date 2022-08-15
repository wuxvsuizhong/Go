package datastruct

import "fmt"

func TestArr2dir(){
	var arr [2][3]int16
	fmt.Println("二维数组，arr:",arr)

	fmt.Printf("arr的类型:%T\n",arr)
	fmt.Printf("arr的地址:%p\n",&arr)
	fmt.Printf("arr[0]的地址:%p\n",&arr[0])
	fmt.Printf("arr[0][0]的地址:%p\n",&arr[0][0])

	arr[0][0] = 11
	arr[0][1] = 12
	arr[1][0] = 13
	arr[1][2] = 14

	fmt.Println("二维数组赋值后，arr:",arr)

	var arr2 [2][2]int16 = [2][2]int16{{10,20},{30,40}}
    fmt.Println("arr2 初始化后，arr2:",arr2)

}

func TravelArr2dir(){
	var arr [3][3]int = [3][3]int{{11,22,33},{44,55,66},{77,88,99}}
	for i := 0;i<len(arr);i++ {
		for j :=0;j<len(arr[i]);j++ {
			fmt.Printf("arr[%v][%v]:%v\t",i,j,arr[i][j])
		}
		fmt.Println()
	}
}

func TravelArr2dir2(){
		fmt.Println("_________________________for key,val := range arr")
	var arr [3][3]int = [3][3]int{{10,20,30},{40,50,60},{70,80,90}}
	for i,val := range arr{
		for j,val2 := range val {
			fmt.Printf("arr[%v][%v]:%v\t",i,j,val2)
		}
		fmt.Println()
	}
}
