package stumanage

import(
	"fmt"
	"os"
)

type student struct{
	name string
	age int
}

//map存储所有学生对象的指针
var allStu map[int64]*student


func showAllStuInfo(){
	for i,val := range allStu {
		fmt.Printf("编号：%d\t姓名：%s\t年龄：%d\n",i,val.name,val.age)
	}
}

func addStu(){
	//创建stu对象
	var stu student 
	fmt.Print("输入学生姓名:")
	fmt.Scanln(&stu.name)
	fmt.Print("输入学生年龄:")
	fmt.Scanln(&stu.age)
	var key int64
	fmt.Print("设定一个编号:")
	fmt.Scanln(&key)
	allStu[key] = &stu
}

func delStu(){
	var delkey int64
	fmt.Print("请输入要删除的序号:")
	fmt.Scanln(&delkey)
	delete(allStu,delkey)
}



func StuMana(){
	allStu = make(map[int64]*student,100)
	for{
		fmt.Println("欢迎进入学员管理系统!")
		fmt.Println(`
			1.查看所有学员信息
			2.新增学员
			3.删除学员
			4.退出
		`)
		fmt.Print("选择:")
		var inputNum int
		fmt.Scanln(&inputNum)
		fmt.Printf("您选择了:%d\n",inputNum)

		switch inputNum{
			case 1:
				showAllStuInfo()
			case 2:
				addStu()
			case 3:
				delStu()
			case 4:
				os.Exit(1)
			default:
				fmt.Println("输入非法!")
		}
	}
}
