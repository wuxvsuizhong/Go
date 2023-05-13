package service

import (
	"customManageSys/model"
	"fmt"
)

//完成对cutomer的操作
//增删改查
type CustomService struct {
	customers   []model.Customer
	customerNum int
}

//初始化创建一个服务实例
func NewCustomerService() *CustomService {
	customerservice := CustomService{
		customers:   []model.Customer{},
		customerNum: 0,
	}
	//id, age int, name, gender, phone, email string
	customerservice.customerNum += 1
	initialOneCustomer := model.NewCuntomer(customerservice.customerNum, 25, "张三", "男", "12345678", "abc.qq.com")
	customerservice.customers = append(customerservice.customers, initialOneCustomer)

	return &customerservice
}

//返回所有客户切片
func (cs *CustomService) AllCustomer() []model.Customer {
	return cs.customers
}

//添加客户到customers切片里
func (cs *CustomService) AddOneCustomer(customer model.Customer) bool {
	cs.customerNum += 1
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

//通过id在切片customers中查找用户,返回用户在切片中的需要，找不到时返回-1
func (cs *CustomService) FindCustomerById(id int) int {
	for i, c := range cs.customers {
		if c.Id == id {
			return i
		}
	}
	// 没找到id
	return -1
}

// 0: 删除成功 -1:删除失败
func (cs *CustomService) RemoveCustomerById(id int) int {
	idx := cs.FindCustomerById(id)
	if idx == -1 {
		return -1
	}
	//从切片customers中删除下标idx的元素
	cs.customers = append(cs.customers[:idx], cs.customers[idx+1:]...)
	return 0
}

func (cs *CustomService) GetOneCustomerById(id int) *model.Customer {
	idx := cs.FindCustomerById(id)
	if idx == -1 {
		return nil
	}
	return &cs.customers[idx]
}

// 0 更新成功 -1 更新失败
func (cs *CustomService) UpdateCustomer(idx int, info model.Customer) int {
	target := &cs.customers[idx]
	var effectInt = func(s *int, d *int) {
		if (*s) != 0 {
			*d = *s
		}
	}

	var effectStr = func(s *string, d *string) {
		if (*s) != "" {
			fmt.Printf("%s 修改为 %s", *d, *s)
			*d = *s
		}
	}

	//target.Age = info.Age
	effectInt(&info.Age, &target.Age)
	//target.Name = info.Name
	effectStr(&info.Name, &target.Name)
	//target.Gender = info.Gender
	effectStr(&info.Gender, &target.Gender)
	//target.Email = info.Email
	effectStr(&info.Email, &target.Email)
	//target.Phone = info.Phone
	effectStr(&info.Phone, &target.Phone)
	//fmt.Println("修改后:", target)
	return 0
}
