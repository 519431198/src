package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 定义连接池
var dbPool *sql.DB

func main() {
	//db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10)/bill")
	//if err != nil {
	//	fmt.Println(errors.New("数据库连接失败"), err)
	//	return
	//}
	//defer db.Close()
	//err = db.Ping()
	//if err != nil {
	//	fmt.Println(err)
	//}
	initDB()
	// 获取连接池中的连接
	db := dbPool
	// 创建了一个5秒超时的 context ,并返回 context 和 取消函数
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 执行 SQL 命令
	// 使用 db.ExecContext() 将 context 传给数据库请求,这会监测 context 是否还活跃,一旦 context 被取消,数据库请求将被关闭。
	rows, err := db.QueryContext(ctx, "select * from `bill_ware_house`")
	//res, err := db.ExecContext(ctx, "select * from `bill_ware_house`")
	if err != nil {
		fmt.Printf("%s跳过\n", err)
	}
	for rows.Next() {

	}
}

// 初始化连接池
func initDB() {
	db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10)/bill")
	if err != nil {
		fmt.Println(errors.New("数据库连接失败"), err)
	}
	// 设置最大连接数
	db.SetMaxOpenConns(20)
	// 设置最大空闲数,最大空闲数一定要小于最大连接数
	db.SetMaxIdleConns(10)
	// 检查连接是否正常
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	dbPool = db
}
