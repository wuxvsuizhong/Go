package view

import (
	"customManageSys/model"
	"customManageSys/service"
	"fmt"
)

type customerView struct {
	op             string
	loop           bool
	customeService *service.CustomService
}

func NewCustomView() *customerView {
	cv := customerView{
		op:             "",
		loop:           true,
		customeService: service.NewCustomerService(),
	}
	return &cv
}

//答应全部客户信息
func (cv *customerView) TravelCustomers() {
	items := cv.customeService.AllCustomer()
	fmt.Println("------------------------客  户  列  表------------------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for _, v := range items {
		fmt.Println(v.GetInfo())
	}
	fmt.Println("-----------------------客 户 列 表 完 成-----------------------")
	fmt.Println()
}

//获取用户输入，构建customer实例，调用service层完成客户信息添加
func (cv *customerView) AddOneCustomer() {
	fmt.Println("------------------------添  加  客  户------------------------")
	customer := model.Customer{}
	fmt.Print("\r输入姓名:")
	fmt.Scanln(&customer.Name)
	fmt.Print("\r输入年龄:")
	fmt.Scanln(&customer.Age)
	fmt.Print("\r输入性别:")
	fmt.Scanln(&customer.Gender)
	fmt.Print("\r输入手机号码:")
	fmt.Scanln(&customer.Phone)
	fmt.Print("\r输入邮箱:")
	fmt.Scanln(&customer.Email)
	if cv.customeService.AddOneCustomer(customer) {
		fmt.Println("\r------------------------添  加  完  成------------------------")
	} else {
		fmt.Println("\r------------------------添  加  失  败------------------------")
	}
}

func (cv *customerView) RemoveCustomerById() {
	fmt.Println("\r------------------------删  除  客  户------------------------")
	fmt.Print("\r请输入要删除客户的ID:")
	id := 0
	fmt.Scanln(&id)

	for {
		fmt.Print("确定要删除吗?y/n:")
		op := ""
		fmt.Scanln(&op)
		if op == "y" || op == "Y" {
			break
		} else {
			return
		}
	}
	if cv.customeService.RemoveCustomerById(id) == -1 {
		fmt.Println("没找到该id对应的用户!")
		fmt.Println("\r------------------------删  除  失  败------------------------")
	} else {
		fmt.Println("\r------------------------删  除  成  功------------------------")
	}
	return
}

func (cv *customerView) exit() {
	fmt.Print("\r确认要退出吗?y/n:")
	op := ""
	for {
		fmt.Scanln(&op)
		if op == "Y" || op == "y" || op == "n" || op == "N" {
			break
		} else {
			fmt.Print("\r输入有误,输入y/n确认是否要退出:")
		}
	}
	if op == "y" || op == "Y" {
		cv.loop = false
	}
}

//返回true 有效修改 false 无效修改
//func chechinput(a interface{}) bool {
//	if a1, ok := a.(*int); ok {
//		for {
//			if n, err := fmt.Scanln(a1); err != nil {
//				fmt.Print("\r输入类型需要是数字,请重新输入:")
//			} else if n == 0 {
//				return false
//			}
//			break
//		}
//	} else if a2, ok := a.(*string); ok {
//		if n, _ := fmt.Scanln(a2); n == 0 {
//			return false
//		}
//	}
//	return true
//}

func (cv *customerView) UpdateById() {
	fmt.Println("\n------------------------修  改  信  息------------------------")
	fmt.Print("请输入要修改的用户ID(-1退出):")
	idx := 0
	fmt.Scanln(&idx)
	if idx == -1 {
		return
	}
	idx = cv.customeService.FindCustomerById(idx)
	if idx != -1 {
		tmpcus := model.Customer{}
		fmt.Print("\r输入姓名:")
		fmt.Scanln(&tmpcus.Name)
		fmt.Print("\r输入年龄:")
		fmt.Scanln(&tmpcus.Age)
		fmt.Print("\r输入性别:")
		fmt.Scanln(&tmpcus.Gender)
		fmt.Print("\r输入手机号码:")
		fmt.Scanln(&tmpcus.Phone)
		fmt.Print("\r输入邮箱:")
		fmt.Scanln(&tmpcus.Email)
		cv.customeService.UpdateCustomer(idx, tmpcus)
		fmt.Println("\n------------------------修  改  完  成------------------------")
	} else {
		fmt.Println("没找到ID对应的客户信息!")
		fmt.Println("\n------------------------修  改  失  败------------------------")
	}
}

func (cv *customerView) MainMenu() {
	for {
		fmt.Println("------------------------客户信息管理软件------------------------")
		fmt.Println("                        1.添 加 客 户                          ")
		fmt.Println("                        2.修 改 客 户                          ")
		fmt.Println("                        3.客 户 列 表                          ")
		fmt.Println("                        4.删 除 客 户                          ")
		fmt.Println("                        5.退      出                          ")
		fmt.Print("\r请选择：")

		fmt.Scanln(&cv.op)
		switch cv.op {
		case "1":
			//fmt.Println("添加客户")
			cv.AddOneCustomer()
			cv.op = "0"
		case "2":
			//fmt.Println("修 改 客 户")
			cv.UpdateById()
			cv.op = "0"
		case "3":
			//fmt.Println("客 户 列 表")
			cv.TravelCustomers()
			cv.op = "0"
		case "4":
			//fmt.Println("删 除 客 户 ")
			cv.RemoveCustomerById()
			cv.op = "0"
		case "5":
			//cv.loop = false
			cv.exit()
			cv.op = "0"
		default:
			fmt.Println("输入有误，请重新输入!")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出客户管理系统!")
}
