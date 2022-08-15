package user

import "fmt"

func GetInput(){
    var age int
    fmt.Println("输入年龄:")
    fmt.Scanln(&age)

    var name string
    fmt.Println("输入姓名:")
    fmt.Scanln(&name)

    var score float32
    fmt.Println("输入成绩:")
    fmt.Scanln(&score)

    var isVIP bool
    fmt.Println("是否VIP:")
    fmt.Scanln(&isVIP)

    fmt.Printf("姓名:%v,年龄:%v,成绩:%v,是否VIP:%v\n",name,age,score,isVIP)

}

func GetInput2(){
	var age2 int
	var name2 string
	var score2 float32
	var isVIP2 bool

	fmt.Println("一次输入姓名，年龄，成绩，是否VIP，空格分割:")
	fmt.Scanf("%d %s %f %t",&age2,&name2,&score2,&isVIP2)
    fmt.Printf("姓名:%v,年龄:%v,成绩:%v,是否VIP:%v\n",name2,age2,score2,isVIP2)
}
