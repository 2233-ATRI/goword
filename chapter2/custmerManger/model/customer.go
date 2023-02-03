package model

import "fmt"

// 声明一个结构体，表面一个客户信息
type Customer struct {
	Id     int64
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式
func NewCustomer(id int64, name string, gender string, Age int, Phone string, Email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}
}

func NewCustomer2(name string, gender string, Age int, Phone string, Email string) Customer {
	return Customer{

		Name:   name,
		Gender: gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}
}

// 返回用户信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
