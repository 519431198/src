package model

import (
	"fmt"
	"github.com/my/repo/web/web_mysql/utils"
)

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

func (user *User) addUser() error {
	// 1.写 sql 语句
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	// 2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常!", err)
	}
	// 3.预编译后执行
	_, err = inStmt.Exec("admin", "123456", "admin@qq.com")
	if err != nil {
		fmt.Println("预编译后执行出现异常", err)
	}
	return err
}

func (user *User) addUser2() error {
	// 1。写 sql 语句
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	// 4.直接执行
	_, err := utils.Db.Exec(sqlStr, "alayman1", "666", "555@qq.com")
	if err != nil {
		fmt.Println("直接执行出现异常!")
	}
	return err
}

func (user *User) query() (int, error) {
	// 1。写 sql 语句
	sqlStr := "SELECT count(*) from users"
	// 4.直接执行
	rows, err := utils.Db.Query(sqlStr)
	var num int
	rows.Scan(&num)
	fmt.Println(num)
	return num, err
}
