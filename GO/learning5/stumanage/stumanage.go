package stumanage

import(
	"fmt"
	"os"
)

func showMenu(){
	fmt.Println(`
		1.查看所有学员信息
		2.添加学员
		3.删除学员
		4.修改学员
		5.退出

		`)
}

var smr studentMgr

func Start(){
	var smr = studentMgr{
		allStudent:make(map[int64]student,100),
	}
	fmt.Println("欢迎来到学员管理系统!")
	for{
		showMenu()
		var choice int
		fmt.Print("请输入选项：")
		fmt.Scanln(&choice)
		fmt.Println("您选择了:",choice)
		switch choice {
		case 1:
			smr.showAllStu()
		case 2:
			smr.addStu()
		case 3:
			smr.delStu()
		case 4:
			smr.editStu()
		case 5:
			os.Exit(1)
		}
	}
}
