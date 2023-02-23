package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("存款金额不正确")
		return
	}
	account.Balance += money
}

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("取款金额不正确")
		return
	}
	account.Balance -= money
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("账户: %v\n余额: %v\n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "Alayman",
		Pwd:       "8888",
		Balance:   10000,
	}
	account.Query("8888")
	account.Deposite(1000, "8888")
	account.Query("8888")
	account.WithDraw(100, "8888")
	account.Query("8888")
}
