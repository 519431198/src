package service

import (
	"go_code/chapter11/customerInfo/model"
)

// CustomerService 完成对 customerService 的操作,包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段,表示当前切片含有多少个客户
	//该字段还可以作为新客户的 id+1
	customerNum int
}

// NewCustomerService 编写一个方法,可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	//调用 model.NewCustomer方法初始化一个客户
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@souhu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// List 返回客户切片
func (that *CustomerService) List() []model.Customer {
	return that.customers
}

func (that *CustomerService) Add(customer model.Customer) bool {
	that.customerNum++
	customer.Id = that.customerNum
	that.customers = append(that.customers, customer)
	return true
}

func (that *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(that.customers); i++ {
		if that.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (that *CustomerService) Delete(id int) bool {
	index := that.FindById(id)
	if index == -1 {
		return false
	}
	that.customers = append(that.customers[:index], that.customers[index+1:]...)
	return true
}

func (that *CustomerService) Update(id int, customer model.Customer) bool {
	index := that.FindById(id)
	if index == -1 {
		return false
	}
	customer.Id = id
	that.customers = append(append(that.customers[:index], customer), that.customers[index+1:]...)
	return true
}
