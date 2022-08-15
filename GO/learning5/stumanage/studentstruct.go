package stumanage
import "fmt"

type student struct{
	id int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) showAllStu(){
	for _,val := range s.allStudent {
		fmt.Printf("学员Id:%d\t学员姓名:%s\n",val.id,val.name)
	}
}

func (s studentMgr) addStu(){
	var newstu student
	fmt.Print("输入学生ID:")
	fmt.Scanln(&newstu.id)
	fmt.Print("输入学员姓名:")
	fmt.Scanln(&newstu.name)

	s.allStudent[newstu.id] = newstu
	fmt.Println("添加成功!")
}

func (s studentMgr) delStu(){
	var stuid int64
	fmt.Print("输入要删除的学员ID:")
	fmt.Scanln(&stuid)
	_,ok := s.allStudent[stuid]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent,stuid)
	fmt.Println("删除完成!")
}

func (s studentMgr) editStu(){
	var stuid int64
	fmt.Print("请输入学员id:")
	fmt.Scanln(&stuid)
	val,ok := s.allStudent[stuid]
	if !ok{
		fmt.Println("查无此人")
		return
	}
	fmt.Println("该学员信息如下：")
	fmt.Printf("学员ID:%d\t姓名:%s\n",val.id,val.name)
	fmt.Print("输入新的姓名:")
	fmt.Scanln(&val.name)
	s.allStudent[stuid] = val
}

