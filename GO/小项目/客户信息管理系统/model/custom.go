package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCuntomer(id, age int, name, gender, phone, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (c Customer) GetInfo() string {
	//编号\t姓名\t性别\t年龄\t电话\t邮箱\
	return fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
