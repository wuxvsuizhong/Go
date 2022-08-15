package datastruct

import "fmt"

func TestInit(){
	var arr1 [3]int = [3]int{22,33,44}
	fmt.Println(arr1)

	var arr2 = [3]int{11,22,33}
	fmt.Println(arr2)

	var arr3 = [...]int{12,13,14,15}
	fmt.Println(arr3)

	var arr4 = [...]int{3:99,0:11,2:66,1:10}
	fmt.Println(arr4)
}
