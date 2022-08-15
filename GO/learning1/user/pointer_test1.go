package user

import "fmt"

func PointerTest(){
	var age int = 20
	fmt.Println(&age)

	var ptr* int = &age
	fmt.Println(ptr)
	fmt.Println("指针ptr变量的地址是:",&ptr)
	fmt.Println("指针ptr指向的值是:",*ptr)

	*ptr += 1
	fmt.Println("指针ptr指向修改后的值是:",*ptr)

	var s1 string = "golang"
	var sptr* string = &s1
	fmt.Printf("字符串s1的地址是%v\n",sptr)
	*sptr += " my test"
	fmt.Printf("修改后的字符串的值:%v\n",*sptr)
}

