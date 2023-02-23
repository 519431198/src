package main

import (
	"fmt"
)

func main() {
	//接收软件控制
	var key string
	//控制软件退出
	var loop = true
	//记录余额
	var balance float64 = 10000
	//记录每次的收支金额
	var money float64
	//记录每次收支的说明
	var note string
	//收支记录详情
	//当有收支时,对字符串进行拼接处理
	var details = ""
	//定义一个 flag
	//var flag = false

	for {
		fmt.Println("\n------------家庭收支记账明细------------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&key)

		switch key {
		case "1":
			if details == "" {
				fmt.Println("当前没有收支明细,来一笔吧")
				break
			}
			//if flag == false {
			//	fmt.Println("当前没有收支明细,来一笔吧")
			//	break
			//}

			fmt.Println("--------------当前收支记录--------------")
			fmt.Println("收支\t账户余额\t收支金额\t说   明")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			//将收入情况,拼接到 details 变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			//flag = true
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足,无法支出!")
				break
			}
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			//将收入情况,拼接到 details 变量
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			//flag = true
		case "4":
			//var choice string
			//fmt.Println("确定要退出吗? y/n")
			//for {
			//	fmt.Scanln(&choice)
			//	if choice == "y" || choice == "n" {
			//		break
			//	}
			//	fmt.Println("输入有误,重新输入 y/n")
			//}
			//if choice == "y" {
			//	loop = false
			//}
			var choice string
			fmt.Println("确定要退出吗? y/n")
			for {
				fmt.Scanln(&choice)
				switch choice {
				case "y":
					loop = false
				case "n":
				default:
					fmt.Println("输入有误,重新输入 y/n")
				}
				break
			}
		default:
			fmt.Println("请输入正确的选项(1-4)")
		}
		if !loop {
			break
		}
	}
	fmt.Println("退出家庭记账软件!")
}
