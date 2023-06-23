package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户:")
	user := &User{}
	//user.addUser()
	//user.addUser2()
	num, _ := user.query()
	fmt.Println(num)
}
