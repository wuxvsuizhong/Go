package main

import "fmt"

// 给结构体绑定的方法中如果有实现了String方法，那么当使用fmt.Println在打印这个方法时就会默认调用它

type Stu struct{
	Name string
	Age int
}

func (s *Stu) String() string{
	str := fmt.Sprintf("Name=[%s]\t Age=[%d]\n",s.Name,s.Age)
	return str
}

func main(){
	s := Stu{"小明",24}

	//直接传递结构体实例，调用的是系统的默认打印方式
	fmt.Println(s)	//{小明 24}
	// 传递结构体指针时调用的是绑定的String方法，因为String绑定使用的是指针
	fmt.Println(&s)	//Name=[小明]      Age=[24]
}