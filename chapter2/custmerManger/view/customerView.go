package main

import (
	"chapter2/custmerManger/model"
	"chapter2/custmerManger/service"
	"fmt"
)

type customerView struct {
	key  string //接受用户输入
	loop bool   //判断是否循环的显示主菜单
	//增加字段customerservice
	CustomerService *service.CustomerService
}

// 显示所有的用户消息
func (this *customerView) list() {
	//切片当中获取信息
	customer := this.CustomerService.List()
	fmt.Println("客户列表")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}

	fmt.Println("客户列表完成")

}
func (this *customerView) mainMenu() {
	for {
		fmt.Println("客户信息管理软件")
		fmt.Println("1.添加客户")
		fmt.Println("2 修改客户")
		fmt.Println("3 删除客户")
		fmt.Println("4 客户列表")
		fmt.Println("5 退出")
		fmt.Println("请选择1-5")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("输入有错误，重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出")
}

// 得到用户的信息，信息构建新的客户，完成添加
func (this *customerView) add() {
	fmt.Println("添加客户")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer
	//系统自动分配id，确保id唯一
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.CustomerService.Add(customer) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.mainMenu()
	customerView.CustomerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
