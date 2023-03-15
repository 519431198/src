package main

import (
	"fmt"
	"github.com/my/repo/go_code/chapter11/customerInfo/model"
	"github.com/my/repo/go_code/chapter11/customerInfo/service"
)

type customerView struct {
	//定义必要字段
	key  string
	loop bool //判断是否循环显示主菜单
	//增加一个字段 CustomerService
	customerService *service.CustomerService
}

func (that *customerView) list() {
	//获取到当前所有的客户信息(在切片中)
	customers := that.customerService.List()
	fmt.Println("------------------客户列表------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("----------------客户列表完成----------------\n\n")
}

// 得到用户的输入信息,构建新的客户,并完成添加
func (that *customerView) add() {
	fmt.Println("------------------添加客户------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机号:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的 customer 实例
	//id 号,让系统自动生成
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if that.customerService.Add(customer) {
		fmt.Println("------------------添加完成------------------")
	} else {
		fmt.Println("------------------添加失败------------------")
	}
}

func (that *customerView) delete() {
	fmt.Println("------------------删除客户------------------")
	fmt.Println("请选择待删除的客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(y/n):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if that.customerService.Delete(id) {
			fmt.Println("------------------删除成功------------------")
		} else {
			fmt.Println("-------------删除失败,id 号不存在-------------")
		}
	}
}

func (that *customerView) update() {
	fmt.Println("------------------修改客户------------------")
	fmt.Println("请输入修改客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机号:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if that.customerService.Update(id, customer) {
		fmt.Println("------------------修改完成------------------")
	} else {
		fmt.Println("------------------修改失败------------------")
	}
}

func (that *customerView) exit() {
	fmt.Println("确认是否退出(y/n):")
	for {
		fmt.Scanln(&that.key)
		if that.key == "y" || that.key == "Y" || that.key == "n" || that.key == "N" {
			break
		}
		fmt.Println("输入有误,请重新输入(y/n)")
	}
	if that.key == "y" || that.key == "Y" {
		that.loop = false
	}
}

// 显示主菜单
func (that *customerView) mainMenu() {
	for {
		fmt.Println("--------------客户信息管理软件--------------")
		fmt.Println("               1 添加客户")
		fmt.Println("               2 修改客户")
		fmt.Println("               3 删除客户")
		fmt.Println("               4 客户列表")
		fmt.Println("               5 退   出")
		fmt.Println("请选择 1-5: ")
		fmt.Scanln(&that.key)
		switch that.key {
		case "1":
			that.add()
		case "2":
			that.update()
		case "3":
			that.delete()
		case "4":
			that.list()
		case "5":
			that.exit()
		default:
			fmt.Println("输入有误,请重新输入!")
		}
		if !that.loop {
			break
		}
	}
	fmt.Println("退出程序")
}

func main() {
	//运行并显示主菜单
	var customerView = customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
