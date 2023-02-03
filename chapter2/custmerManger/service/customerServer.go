package service

import "chapter2/custmerManger/model"

// 该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	customer []model.Customer
	//声明一个字段，表明当前切片含有多少个客户
	customerNum int64
	//该字段后面可以作为新客户Id

}

// 方法，可以返回*customerserice
func NewCustomerService() *CustomerService {
	//方便看到切片进行用户的初始化
	CustomerService := &CustomerService{}
	CustomerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 19, "112", "zs@shscoab.com")
	CustomerService.customer = append(CustomerService.customer, customer)
	return CustomerService
}

//返回客户的切片

func (this *CustomerService) List() []model.Customer {
	return this.customer
}

// 添加客户到costomer切片当中
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customer = append(this.customer, customer) //添加进去
	return true
}
