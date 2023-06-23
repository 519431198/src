package utils

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// Res 定义结果字段类型
type Res struct {
	CustomerName string
	Charge6      string
	Charge60     string
}

// Customer 解析配置文件
type Customer struct {
	Phones []string
}

type Customers struct {
	Customers map[string]Customer
}

// DbPool 定义连接池
var DbPool *sql.DB

// Ph 定义结构体
type Ph struct {
}

// Test 定义接口
type Test interface {
	InitDB()
	NumDigits(s string) bool
}

// InitDB 初始化连接池
func (ph *Ph) InitDB() {
	// 设置连接池的最大连接数和最大空闲连接数
	db, err := sql.Open("mysql", "dandan:dwy123@tcp(39.102.237.93:12306)/bill")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// 检查连接池是否正常工作
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	DbPool = db
}

// NumDigits 判断号码是否为 11 位
func (ph *Ph) NumDigits(s string) bool {
	// 编写正则表达式
	rule := `^\d{11}$`
	// 使用正则表达式
	r := regexp.MustCompile(rule)
	// 返回判断结果
	return r.MatchString(s)
}

// GetSql 获取 sql 字符串
func GetSql(Customer Customer, test Test) string {
	var phone string
	for _, v := range Customer.Phones {
		// 移除号码两端的空格
		v = strings.TrimSpace(v)
		// 判断号码是否为 11 位
		if !test.NumDigits(v) {
			fmt.Printf("%s 号码格式不对\n", v)
			continue
		}
		// 拼接字符串
		phone += fmt.Sprintf("'%s',", v)
	}
	phone = strings.TrimSuffix(phone, ",")
	// 拼接查询表名
	form := fmt.Sprintf("bill_ware_house_" + time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	// 写 sql 语句
	//sqlStr := fmt.Sprintf("SELECT ? as customer_name,sum(charge6),sum(charge60) FROM `bill_ware_house_2023-06-21` WHERE phone_number_x IN (%s) AND charge60 != '0'", phone)
	sqlStr := fmt.Sprintf("SELECT ? as customer_name,sum(charge6),sum(charge60) FROM `%s` WHERE phone_number_x IN (%s) AND charge60 != '0'", form, phone)
	return sqlStr
}

// Exec 执行 sql 查询语句
func Exec(sqlStr, name string) string {
	// 预编译
	inStmt, err := DbPool.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常!", err)
	}
	// 预编译后执行
	row := inStmt.QueryRow(name)
	res := Res{}
	// 赋值给对应变量
	err = row.Scan(&res.CustomerName, &res.Charge6, &res.Charge60)
	if res.Charge6 == "" && res.Charge60 == "" {
		res.Charge6 = "0"
		res.Charge60 = "0"
	} else if err != nil {
		fmt.Println("rows.Scan 赋值失败:", err)
	}
	data := fmt.Sprintf("%s,%v,%v\n", res.CustomerName, res.Charge6, res.Charge60)
	return data
}

// WriteData 查询数据写入文件
func WriteData(str, dataPath string) {
	file, err := os.OpenFile(dataPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("数据写入失败", err)
	}
	err = file.Sync()
	if err != nil {
		fmt.Println("刷新写入失败", err)
	}
}
