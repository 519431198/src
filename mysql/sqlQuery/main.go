package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/my/repo/mysql/sqlQuery/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var test utils.Test
	test = &utils.Ph{}

	// 获取该可执行文件所在的目录路径
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}

	configPath := filepath.Join(binPath, "config.yaml")
	resDataPath := filepath.Join(binPath, "sqlQuery.csv")

	// 创建 config 实例
	var config utils.Customers
	// 读取配置文件
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("os.Open() err:", err)
	}
	// 创建一个读取器
	reader := bufio.NewReader(file)
	// 解析配置文件到结构体
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("decoder.Decode() err:", err)
	}
	// 初始化数据库连接池
	test.InitDB()
	// 写入查询日期
	file, err = os.OpenFile(resDataPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("时间,%s\n客户名,6秒数,分钟数\n", time.Now().AddDate(0, 0, -1).Format("2006-01-02")))
	// 查询数据
	for name, Customer := range config.Customers {
		sqlStr := utils.GetSql(Customer, test)
		data := utils.Exec(sqlStr, name)
		utils.WriteData(data, resDataPath)
	}
}
