package main

import "fmt"

func main() {
	//声明一个变量保存输入的选项
	key := ""
	//声明一个变量,控制退出for循环
	loop := true
	//显示主菜单
	banlance := 10000.0
	//账户的余额
	money := 0.0
	//每次收支的金额
	note := ""
	//每次收支的说明
	//收支的详细情况
	details := "收支\t账户金额\t收支金额\t说明"
	for {
		fmt.Println("家庭收支记账软件")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")
		fmt.Print("请输入（1-4）:")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("当前收支明细记录")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			banlance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", banlance, money, note)
		case "3":
			fmt.Println("登记支出：")
			fmt.Scanln(&money)
			if money > banlance {
				fmt.Println("金额不足")
				break
			}
			banlance -= money
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", banlance, money, note)
		case "4":
			fmt.Println("是否确定退出y/n")
			chice := ""
			for {
				fmt.Scanln(&chice)
				if chice == "y" || chice == "n" {
					break
				} else {
					fmt.Println("输入错误")
				}
			}
			if chice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正常选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出软件")
}
