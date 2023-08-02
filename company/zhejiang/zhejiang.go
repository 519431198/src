package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"time"
)

var db *gorm.DB

type BillHn struct {
	Id                    uint   `gorm:"column:id;size:20"`
	PhoneNumberA          string `gorm:"column:phone_number_a;size:50"`
	PhoneNumberX          string `gorm:"column:phone_number_x;size:50"`
	PhoneNumberB          string `gorm:"column:phone_number_b;size:50"`
	PhoneNumberC          string `gorm:"column:phone_number_c;size:50"`
	PhoneNumberY          string `gorm:"column:phone_number_y;size:50"`
	Dgts                  string `gorm:"column:dgts;size:50"`
	BindingId             string `gorm:"column:binding_id;size:50"`
	CallType              string `gorm:"column:call_type;size:50"`
	CallResult            string `gorm:"column:call_result;size:50"`
	CallTime              string `gorm:"column:call_time;size:50"`
	RingingTime           string `gorm:"column:ringing_time;size:50"`
	StartTime             string `gorm:"column:start_time;size:50"`
	ReleaseTime           string `gorm:"column:release_time;size:50"`
	CallId                string `gorm:"column:call_id;size:50"`
	ReleaseDirection      string `gorm:"column:release_direction;size:50"`
	ReleaseCause          string `gorm:"column:release_cause;size:50"`
	CallRecording         string `gorm:"column:call_recording;size:50"`
	RecordingUrl          string `gorm:"column:recording_url;size:256"`
	RecordingMode         string `gorm:"column:recording_mode;size:50"`
	TransferPhoneNumber   string `gorm:"column:transfer_phone_number;size:50"`
	TransferReason        string `gorm:"column:transfer_reason;size:50"`
	Ability               string `gorm:"column:ability;size:50"`
	CallRecognitionResult string `gorm:"column:call_recognition_result;size:50"`
	AdditionalData        string `gorm:"column:additional_data;size:50"`
}

// Customer 解析配置文件
type Customer struct {
	Phones string
}
type Customers struct {
	Customers map[string]Customer
}

//var mysqlLogger logger.Interface

// 初始化数据库
func init() {
	username := "dandan"
	password := "dwy123"
	host := "39.102.237.93"
	port := "12306"
	timeout := "10s"

	// 日志
	//mysqlLogger = logger.Default.LogMode(logger.Info)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/hn_bill?charset=utf8mb4&parseTime=True"+
		"&loc=Local&timeout=%s", username, password, host, port, timeout)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	db = dbClient
}

func main() {
	now := time.Now()
	// 获取前一天时间
	oldTime := now.AddDate(0, 0, -1).Format("2006-01-02")

	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	// 获取当前路径下文件
	//resDataPath := filepath.Join(fmt.Sprintf("%s/%s.csv", binPath, "hnzj_bill"))
	resDataPath := "/Users/wangyi/goProject/project_1/src/company/zhejiang/" + oldTime + ".csv"
	fmt.Println(resDataPath)
	resFile, err := os.OpenFile(resDataPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer resFile.Close()

	// 日志打印,测试用
	//db = db.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	var billHn BillHn

	// 创建结构体实例
	var config Customers
	// 获取当前路径下配置文件
	configPath := binPath + "/config.yaml"
	// 打开配置文件,创建文件句柄
	//file, err := os.Open("/Users/wangyi/goProject/project_1/src/company/zhejiang/config.yaml")
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("文件读取失败: ", err)
	}
	defer file.Close()
	// 创建读取器
	reader := bufio.NewReader(file)
	// 读取配置文件内容
	decoder := yaml.NewDecoder(reader)
	// 反序列化到结构体
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("解析配置文件失败", err)
	}
	// 拼接定义前一天表名
	table := "bill_hainan_" + now.AddDate(0, 0, -1).Format("20060102")
	for name, customer := range config.Customers {
		// 定义 where 条件
		query := fmt.Sprintf("phone_number_x in (%s)", customer.Phones)
		var count int64
		// 查询数据库
		db.Model(&billHn).Table(table).Where(query).Select("sum(ceil( time_to_sec( timediff(release_time, start_time) ) / 60 ))").Scan(&count)
		_, err = resFile.WriteString(fmt.Sprintf("%s,%s,%d\n", oldTime, name, count))
		if err != nil {
			fmt.Println("数据写入失败", err)
		}
		err = file.Sync()
		if err != nil {
			fmt.Println("刷新写入失败", err)
		}
	}
}
